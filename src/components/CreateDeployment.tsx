import {
  Box,
  Button,
  Input,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  Text,
  Icon,
  Stack,
  useTheme,
  FormControl,
  FormLabel,
  Heading,
  Textarea,
  Flex,
} from '@chakra-ui/react'
import { useToast } from '@liftedinit/ui'
import React, { useCallback, useEffect, useState } from 'react'
import { useDropzone } from 'react-dropzone'
import { BiCheck, BiSolidCloudUpload } from 'react-icons/bi'
import { IoWarning } from 'react-icons/io5'
import { FaExclamation } from 'react-icons/fa'
import {
  Deployment,
  fileToArrayBuffer,
  handleDeploymentCreationSuccess,
  useCreateDeployment,
  useUpdateDeployment,
} from '../features/deployments'
import {
  DeploymentTypes,
  DeploymentSource,
  Archive,
  WebDeployInfo,
  WebDeployParams,
} from '@liftedinit/many-js'
import { useAccountsStore } from '../features/accounts'
import { generateRandomString } from '../shared'
import Info from './Info'

const MAX_FILE_SIZE_BYTES = 5242546 // This is the limit supported by the backend, including the envelope and header overhead
const SITE_NAME_MAX_LENGTH = 12
const SITE_DESCRIPTION_MAX_LENGTH = 500
const TRANSACTION_MEMO_MAX_LENGTH = 500
const DOMAIN_MAX_LENGTH = 253

interface FormState {
  siteName: string
  siteDescription?: string
  zipFile: File | undefined
  transactionMemo?: string
  domain?: string
}

const initialFormState = {
  siteName: generateRandomString(12),
  siteDescription: '',
  zipFile: undefined,
  transactionMemo: '', // Must be a controlled element
  domain: '', // Must be a controlled element
}

type CreateDeploymentProps = {
  onClose: () => void
  isOpen: boolean
  deployments: Deployment[]
  activeDeploymentUuid: string
  isRedeploying: boolean
  setIsRedeploying: (value: boolean) => void
}

type Lengths = {
  siteName: number
  siteDescription: number
  transactionMemo: number
  domain: number
}

export default function CreateDeployment({
  onClose,
  isOpen,
  deployments,
  activeDeploymentUuid,
  isRedeploying,
  setIsRedeploying,
}: CreateDeploymentProps) {
  const theme = useTheme()
  const toast = useToast()
  const { acceptedFiles, fileRejections, getRootProps, getInputProps } =
    useDropzone({
      accept: {
        'application/zip': [],
      },
      maxFiles: 1,
      maxSize: MAX_FILE_SIZE_BYTES,
      onDrop: files => {
        setError(null)
        setFormState({
          ...formState,
          zipFile: files[0],
        })
      },
    })
  const [formState, setFormState] = useState<FormState>(initialFormState)
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [isComplete, setIsComplete] = useState(false)
  const [error, setError] = useState<Error | null>(null)
  const isDisabled = !formState.zipFile || !formState.siteDescription?.length
  const [lengths, setLengths] = useState<Lengths>({
    siteName: 0,
    siteDescription: 0,
    transactionMemo: 0,
    domain: 0,
  })
  const account = useAccountsStore(s => s.byId.get(s.activeId))
  const createDeploymentMutation = useCreateDeployment(account?.address)
  const updateDeploymentMutation = useUpdateDeployment(account?.address)

  const remainingChars = (currentLength: number, maxLength: number) =>
    `${currentLength}/${maxLength}`
  const findActiveDeployment = () =>
    deployments.find(deployment => deployment.uuid === activeDeploymentUuid)

  const handleInputChange = useCallback(
    (event: React.FormEvent) => {
      const { name, value } = event.target as HTMLInputElement
      if (Object.keys(lengths).includes(name)) {
        setLengths(prev => ({ ...prev, [name]: value.length }))
        setFormState(prev => ({ ...prev, [name]: value }))
      }
    },
    [lengths],
  )

  useEffect(() => {
    if (!isOpen) return

    const foundDeployment = findActiveDeployment()
    if (foundDeployment) {
      setFormState(prevState => ({
        ...prevState,
        siteName: foundDeployment.siteName,
        siteDescription: foundDeployment.siteDescription,
        transactionMemo: foundDeployment.transactionMemo,
        zipFile: undefined,
        domain: foundDeployment.domain,
      }))
    }
  }, [isOpen]) // eslint-disable-line react-hooks/exhaustive-deps

  const handleMutation = async (
    data: WebDeployParams,
    isRedeploying: boolean,
    mutations: { create: any; update: any },
    callbacks: { onSuccess: any; onError: any },
  ) => {
    const mutation = isRedeploying ? mutations.update : mutations.create

    mutation.mutate(data, {
      onSuccess: callbacks.onSuccess,
      onError: callbacks.onError,
    })
  }

  const handleSuccess = useCallback(
    (returnedData: WebDeployInfo) => {
      handleDeploymentCreationSuccess(setError, {
        returnedData,
        deployments,
        activeDeploymentUuid,
        setIsSubmitting,
        setIsComplete,
        setIsRedeploying,
      })
    },
    [deployments, activeDeploymentUuid, setIsRedeploying],
  )

  const handleError = useCallback(
    (error: Error) => {
      setIsSubmitting(false)
      setError(error)
      toast({
        title: 'Error',
        description: error.message,
        status: 'error',
      })
      setIsRedeploying(false)
    },
    [setIsRedeploying, toast],
  )

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault()
    setIsSubmitting(true)

    if (formState.zipFile) {
      let arrayBuffer: ArrayBuffer
      try {
        arrayBuffer = await fileToArrayBuffer(formState.zipFile)
      } catch (error) {
        setIsSubmitting(false)
        setError(error as Error)
        return
      }
      const payload: Archive = [0, new Map().set(0, arrayBuffer)]
      const deploymentData = {
        owner: account?.address,
        siteName: formState.siteName,
        siteDescription: formState.siteDescription,
        deploymentSource: {
          type: DeploymentTypes.Archive,
          payload,
        } as DeploymentSource,
        memo: formState.transactionMemo ? [formState.transactionMemo] : [''],
        domain: formState.domain,
      }

      await handleMutation(
        deploymentData,
        isRedeploying,
        { create: createDeploymentMutation, update: updateDeploymentMutation },
        {
          onSuccess: handleSuccess,
          onError: handleError,
        },
      )
    }
  }

  const acceptedFileItems = acceptedFiles.map(file => {
    return `${(file as any).path} - ${(file.size / 1024 / 1024).toFixed(1)} MB`
  })

  const handleClose = () => {
    onClose()
    setIsComplete(false)
    setError(null)
    setFormState({ ...initialFormState, siteName: generateRandomString(12) })
    setIsRedeploying(false)
  }

  return (
    <Modal isOpen={isOpen} onClose={handleClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Create Deployment</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          {isComplete ? (
            <Stack spacing={5}>
              <Box display="flex" justifyContent="center" justifyItems="center">
                <Box
                  as="span"
                  w={35}
                  h={35}
                  display="flex"
                  justifyContent="center"
                  alignItems="center"
                  color={theme.colors.white}
                  rounded="full"
                  bg={error ? theme.colors.red[400] : theme.colors.green[400]}
                  mr={2}
                >
                  {error ? (
                    <Icon as={FaExclamation} w={25} h={25} />
                  ) : (
                    <Icon as={BiCheck} w={35} h={35} />
                  )}
                </Box>
              </Box>
              <Box display="flex" justifyContent="center" justifyItems="center">
                <Heading>{error ? 'Error' : 'Success!'}</Heading>
              </Box>
              <Text align="center">
                {error
                  ? 'Something went wrong. Please try again.'
                  : 'Your upload is complete.'}
              </Text>
              <Box
                display="flex"
                justifyContent="center"
                justifyItems="center"
                mt={10}
              >
                <Button
                  type="submit"
                  isLoading={isSubmitting}
                  onClick={handleClose}
                  px={10}
                >
                  Close
                </Button>
              </Box>
            </Stack>
          ) : (
            <form onSubmit={handleSubmit}>
              <Stack spacing={5}>
                <FormControl>
                  <Flex>
                    <FormLabel htmlFor="siteName">Site Name</FormLabel>
                    <Box ml="auto">
                      <Info id="siteName" />
                    </Box>
                  </Flex>
                  <Input
                    id="siteName"
                    name="siteName"
                    placeholder=""
                    size="lg"
                    value={formState.siteName}
                    onChange={handleInputChange}
                    borderColor={theme.colors.gray[400]}
                    maxLength={SITE_NAME_MAX_LENGTH}
                    disabled={true}
                  />
                </FormControl>
                <FormControl isRequired>
                  <Flex align="center">
                    <FormLabel htmlFor="siteDescription">
                      Site Description
                    </FormLabel>
                    <Box ml="auto">
                      <Info id="siteDescription" />
                    </Box>
                  </Flex>
                  <Textarea
                    id="siteDescription"
                    name="siteDescription"
                    placeholder=""
                    size="lg"
                    value={formState.siteDescription}
                    onChange={handleInputChange}
                    borderColor={theme.colors.gray[400]}
                    maxLength={SITE_DESCRIPTION_MAX_LENGTH}
                  />
                  <Text fontSize="sm" color="gray.500">
                    {remainingChars(
                      lengths.siteDescription,
                      SITE_DESCRIPTION_MAX_LENGTH,
                    )}
                  </Text>
                </FormControl>

                <FormControl>
                  <Flex align="center">
                    <FormLabel htmlFor="transactionMemo">
                      Transaction Memo
                    </FormLabel>
                    <Box ml="auto">
                      <Info id="transactionMemo" />
                    </Box>
                  </Flex>
                  <Input
                    id="transactionMemo"
                    name="transactionMemo"
                    placeholder=""
                    size="lg"
                    value={formState.transactionMemo ?? ['']}
                    onChange={handleInputChange}
                    borderColor={theme.colors.gray[400]}
                    maxLength={TRANSACTION_MEMO_MAX_LENGTH}
                  />
                  <Text fontSize="sm" color="gray.500">
                    {remainingChars(
                      lengths.transactionMemo,
                      TRANSACTION_MEMO_MAX_LENGTH,
                    )}
                  </Text>
                </FormControl>

                <FormControl>
                  <Flex align="center">
                    <FormLabel htmlFor="domain">Custom domain</FormLabel>
                    <Box ml="auto">
                      <Info id="domain" />
                    </Box>
                  </Flex>
                  <Input
                    id="domain"
                    name="domain"
                    placeholder=""
                    size="lg"
                    value={formState.domain}
                    onChange={handleInputChange}
                    borderColor={theme.colors.gray[400]}
                    maxLength={DOMAIN_MAX_LENGTH}
                  />
                  <Text fontSize="sm" color="gray.500">
                    {remainingChars(lengths.domain, DOMAIN_MAX_LENGTH)}
                  </Text>
                </FormControl>

                <FormControl isRequired>
                  <Flex align="center">
                    <FormLabel htmlFor="transactionMemo">
                      Deployment File
                    </FormLabel>
                    <Box ml="auto">
                      <Info id="deploymentFile" />
                    </Box>
                  </Flex>
                  <Box
                    {...getRootProps()}
                    sx={{
                      display: 'flex',
                      flexDirection: 'column',
                      alignItems: 'center',
                      justifyContent: 'center',
                      height: '100px',
                      padding: 4,
                      borderWidth: 1,
                      borderRadius: 6,
                      borderColor: theme.colors.gray[400],
                      backgroundColor: theme.colors.gray[50],
                      color: theme.colors.gray[500],
                      cursor: 'pointer',
                    }}
                  >
                    <input {...getInputProps()} name="zipFile" />
                    <Icon
                      as={BiSolidCloudUpload}
                      w={30}
                      h={30}
                      color={theme.colors.gray[500]}
                      mb={3}
                    />
                    <Text align="center">
                      Drag a zip file here, or click to select a file
                    </Text>
                  </Box>
                  {formState.zipFile ? (
                    <Text
                      display="flex"
                      mt={4}
                      alignItems="center"
                      justifyContent="center"
                    >
                      <Box
                        as="span"
                        w={5}
                        h={5}
                        display="flex"
                        justifyContent="center"
                        color={theme.colors.white}
                        rounded="full"
                        bg={theme.colors.green[400]}
                        mr={2}
                      >
                        <Icon as={BiCheck} w={5} h={5} />
                      </Box>
                      {acceptedFileItems}
                    </Text>
                  ) : null}
                  {fileRejections.length > 0 && (
                    <Box mt={4}>
                      {fileRejections.map(rejection => (
                        <Box
                          key={rejection.file.name}
                          color={theme.colors.red[600]}
                          display="flex"
                          alignItems="center"
                        >
                          <Icon as={IoWarning} w={6} h={6} mr={2} />
                          {rejection.errors.map(error => (
                            <Text key={error.code}>{error.message} </Text>
                          ))}
                        </Box>
                      ))}
                    </Box>
                  )}
                </FormControl>
              </Stack>
              <Box
                display="flex"
                justifyContent="center"
                justifyItems="center"
                mt={10}
              >
                <Button
                  type="submit"
                  isLoading={isSubmitting}
                  disabled={isDisabled}
                  px={10}
                >
                  Save
                </Button>
              </Box>
            </form>
          )}
        </ModalBody>

        <ModalFooter></ModalFooter>
      </ModalContent>
    </Modal>
  )
}
