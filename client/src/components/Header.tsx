import { AppBar, Toolbar, Typography, Button, Container } from '@mui/material';
import { makeStyles } from '@mui/styles';

const useStyles = makeStyles({
  title: {
    flexGrow: 1,
  },
  link: {
    color: '#fff',
    textDecoration: 'none',
  },
  appBar: {
    zIndex: 1400, // Ensure the AppBar is above other components like the Drawer
    width: '100%', // Make sure the AppBar takes the full width
  },
});

function Header() {
  const classes = useStyles();

  return (
    <AppBar position="fixed" className={classes.appBar}>
      <Container maxWidth={false}>
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            My Application
          </Typography>
          <Button color="inherit">
            <a href="/" className={classes.link}>Home</a>
          </Button>
          <Button color="inherit">
            <a href="/deployments" className={classes.link}>Deployments</a>
          </Button>
          {/* Add more links as needed */}
        </Toolbar>
      </Container>
    </AppBar>
  );
}

export default Header;
