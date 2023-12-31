import React from 'react'
import {
  Box,
  Button,
  Container,
  Grid,
  GridItem,
  Image,
  Icon,
  IconButton,
  useDisclosure,
  Stack,
  Link as ChakraLink,
  useTheme,
  Popover,
  PopoverTrigger,
  PopoverContent,
  PopoverBody,
  Text,
} from '@chakra-ui/react'
import { Link as ReactRouterLink } from 'react-router-dom'
import { IoMdMenu, IoMdClose } from 'react-icons/io'
import { BiCheck, BiCopy, BiSolidUser } from 'react-icons/bi'
import { AnonymousIdentity } from '@liftedinit/many-js'

import { useAccountsStore } from '../features/accounts'

import logo from '../assets/logo-black.png'

interface Props {
  children: React.ReactNode
  href?: string
  onClick: () => void
  sx?: any
}

const NavLink = (props: Props) => {
  const { children, href, onClick, sx } = props

  if (!href) {
    return (
      <Box
        px={2}
        py={1}
        onClick={() => onClick()}
        sx={{ ...sx, cursor: 'pointer' }}
      >
        {children}
      </Box>
    )
  }

  return (
    <Box as={ReactRouterLink} px={2} py={1} to={href} onClick={() => onClick()}>
      {children}
    </Box>
  )
}

export default function Header({ onAddModalOpen }: any) {
  const account = useAccountsStore(s => s.byId.get(s.activeId))
  const isAnonymous = account?.identity instanceof AnonymousIdentity
  const [copied, setCopied] = React.useState<boolean>(false)

  const {
    isOpen: menuOpen,
    onOpen: onMenuOpen,
    onClose: onMenuClose,
  } = useDisclosure()

  const {
    isOpen: accountOpen,
    onOpen: onAccountOpen,
    onClose: onAccountClose,
  } = useDisclosure()

  const { deleteAccount, accounts } = useAccountsStore(
    ({ byId, deleteAccount }) => ({
      deleteAccount,
      accounts: Array.from(byId),
    }),
  )

  const handleLogin = () => {
    onAddModalOpen()
  }

  const handleLogout = () => {
    deleteAccount(accounts[1][0])
  }

  const theme = useTheme()

  return (
    <Container
      maxWidth="100%"
      sx={{ borderBottom: `1px solid ${theme.colors.gray[300]}` }}
      data-testid="header"
    >
      <Container maxW="4xl">
        <Grid templateColumns={`repeat(3, 1fr)`}>
          <GridItem
            colSpan={1}
            display={'flex'}
            alignItems={'center'}
            justifyContent={'flex-start'}
            py={2}
          >
            {!isAnonymous && (
              <Box
                sx={{
                  '& section:focus': { boxShadow: 'none' },
                }}
              >
                <Popover
                  isOpen={menuOpen}
                  onClose={onMenuClose}
                  placement="bottom-start"
                >
                  <PopoverTrigger>
                    <IconButton
                      size="lg"
                      icon={menuOpen ? <IoMdClose /> : <IoMdMenu />}
                      aria-label="Open Menu"
                      onClick={menuOpen ? onMenuClose : onMenuOpen}
                      bg={theme.colors.white}
                      sx={{
                        display: { base: 'flex', md: 'none' },
                        justifyContent: 'center',
                        ml: -5,
                      }}
                    />
                  </PopoverTrigger>
                  <PopoverContent>
                    <PopoverBody
                      sx={{ boxShadow: '0 2px 3px 1px rgba(0, 0, 0, 0.06)' }}
                      p={4}
                    >
                      <Stack as="nav" spacing={4}>
                        <NavLink href="/dashboard" onClick={onMenuClose}>
                          Dashboard
                        </NavLink>
                      </Stack>
                    </PopoverBody>
                  </PopoverContent>
                </Popover>
              </Box>
            )}
          </GridItem>
          <GridItem
            colSpan={1}
            display={'flex'}
            alignItems={'center'}
            justifyContent={'center'}
            py={2}
          >
            <ChakraLink as={ReactRouterLink} to="/" p={2}>
              <Image h={4} src={logo} sx={{ minWidth: '200px' }} />
            </ChakraLink>
          </GridItem>
          <GridItem
            colSpan={1}
            display={'flex'}
            alignItems={'center'}
            justifyContent={'flex-end'}
            py={2}
          >
            {!isAnonymous && (
              <Box
                display={{ base: 'none', md: 'flex' }}
                as={ReactRouterLink}
                to="/dashboard"
                sx={{ fontWeight: 'bold', mr: 3 }}
              >
                Dashboard
              </Box>
            )}

            <Box
              sx={{
                '& section:focus': { boxShadow: 'none' },
              }}
            >
              <Popover
                isOpen={accountOpen}
                onClose={onAccountClose}
                placement="bottom-start"
              >
                <PopoverTrigger>
                  <IconButton
                    size="lg"
                    icon={<BiSolidUser />}
                    aria-label="Account"
                    onClick={() => {
                      if (isAnonymous) {
                        handleLogin()
                        return
                      }
                      if (accountOpen) {
                        onAccountClose()
                      } else {
                        onAccountOpen()
                      }
                    }}
                    bg={theme.colors.white}
                    mr={-4}
                  />
                </PopoverTrigger>
                <PopoverContent>
                  <PopoverBody
                    sx={{ boxShadow: '0 2px 3px 1px rgba(0, 0, 0, 0.06)' }}
                    p={4}
                  >
                    <Stack as="nav" spacing={4}>
                      <>
                        <Text fontWeight={'bold'}>{account?.name}</Text>
                        <Button
                          variant="outline"
                          sx={{
                            background: copied
                              ? theme.colors.gray[100]
                              : 'none',
                          }}
                          onClick={() => {
                            setCopied(true)
                            setTimeout(() => setCopied(false), 1000)
                            navigator.clipboard.writeText(
                              account?.address.toString() || '',
                            )
                          }}
                        >
                          <Box
                            sx={{
                              pr: 22,
                            }}
                          >
                            {account?.address.toString().slice(0, 4)}...
                            {account?.address.toString().slice(-4)}
                          </Box>
                          <Icon as={copied ? BiCheck : BiCopy} w={5} h={5} />
                        </Button>
                        <Button
                          variant="outline"
                          onClick={() => {
                            handleLogout()
                            onAccountClose()
                          }}
                        >
                          Log Out
                        </Button>
                      </>
                    </Stack>
                  </PopoverBody>
                </PopoverContent>
              </Popover>
            </Box>
          </GridItem>
        </Grid>
      </Container>
    </Container>
  )
}
