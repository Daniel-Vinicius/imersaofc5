import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import StoreIcon from '@mui/icons-material/Store'

export const Navbar = () => {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <StoreIcon />
          <Typography variant="h6" component="h1" sx={{ flexGrow: 1 }}>
            Fincycle
          </Typography>
        </Toolbar>
      </AppBar>
    </Box>
  );
}