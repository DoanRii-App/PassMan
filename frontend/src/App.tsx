import { useState } from 'react';
import ReactDOM from 'react-dom/client';
import Darkmode from './components/Darkmode';
import Logo from './assets/passmanlogo.png';
import Logodark from './assets/passmanlogo-dark.png';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import {
  MantineProvider,
  ColorSchemeProvider,
  ColorScheme,
  AppShell,
  Navbar,
  Header,
  Footer,
  Aside,
  Text,
  MediaQuery,
  Burger,
  useMantineTheme,
  Group,
  Button,
  Box,
  UnstyledButton,
  Avatar
} from '@mantine/core';
import {
  IconLogout,
  IconChevronRight,
  IconChevronLeft
} from '@tabler/icons';
import NaviBar from './components/NaviBar';
import Login from './pages/Login';

function App() {
  const [colorScheme, setColorScheme] = useState<ColorScheme>('light');
  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));
  const theme = useMantineTheme();
  const [opened, setOpened] = useState(false);

  return (
    <ColorSchemeProvider
      colorScheme={colorScheme}
      toggleColorScheme={toggleColorScheme}
    >
      <MantineProvider
        theme={{ colorScheme }}
        withGlobalStyles
        withNormalizeCSS
      >
        <AppShell
          navbarOffsetBreakpoint="sm"
          asideOffsetBreakpoint="sm"
          navbar={ 
            <Navbar
              p="md"
              hiddenBreakpoint="sm"
              hidden={!opened}
              width={{ sm: 200, lg: 300 }}
            ><NaviBar /></Navbar>}
          header={
            <Header height={70} p="md">
              <Group sx={{ height: '100%' }} position="apart">
                <Group>
                  <MediaQuery largerThan="sm" styles={{ display: 'none' }}>
                    <Burger
                      opened={opened}
                      onClick={() => setOpened((o) => !o)}
                      size="sm"
                      color={theme.colors.gray[6]}
                      mr="xl"
                    />
                  </MediaQuery>
                  {colorScheme === 'dark' ? (
                    <img
                      style={{ width: '150px', display: 'flex' }}
                      src={Logodark}
                    ></img>
                  ) : (
                    <img
                      style={{ width: '150px', display: 'flex' }}
                      src={Logo}
                    ></img>
                  )}
                </Group>
                  <Darkmode />
              </Group>
            </Header>
          }
        >
          <BrowserRouter>
          <Routes>
            <Route path="/" element={<Login />} />
          </Routes>
        </BrowserRouter>
        </AppShell>
      </MantineProvider>
    </ColorSchemeProvider>
  );
}

export default App;
