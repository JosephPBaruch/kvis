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
            Kvis
          </Typography>
          <Button color="inherit">
            <a href="/" className={classes.link}>Home</a>
          </Button>
          <Button color="inherit">
            <a href="/deployments" className={classes.link}>Deployments</a>
          </Button>
          <Button color="inherit">
            <a href="/pods" className={classes.link}>Pods</a>
          </Button>
          <Button color="inherit">
            <a href="/services" className={classes.link}>Services</a>
          </Button>
          <Button color="inherit">
            <a href="/configmaps" className={classes.link}>ConfigMaps</a>
          </Button>
          <Button color="inherit">
            <a href="/ingresses" className={classes.link}>Ingresses</a>
          </Button>
          <Button color="inherit">
            <a href="/pvcs" className={classes.link}>PVCs</a>
          </Button>
          <Button color="inherit">
            <a href="/endpoints" className={classes.link}>Endpoints</a>
          </Button>
          <Button color="inherit">
            <a href="/nodes" className={classes.link}>Nodes</a>
          </Button>
          <Button color="inherit">
            <a href="/namespaces" className={classes.link}>Namespaces</a>
          </Button>
        </Toolbar>
      </Container>
    </AppBar>
  );
}

export default Header;
