import { useState, useEffect } from 'react';
import { Drawer, List, ListItemButton, ListItemText, Button } from "@mui/material";
import { makeStyles } from '@mui/styles';
import Header from './components/Header';
import deploymentsList from './utils/http';
import { Deployment } from './types/deployment';

const useStyles = makeStyles({
  root: {
    padding: '20px',
    marginTop: '64px', // Adjust for the height of the AppBar
    transition: 'margin-left 0.3s', // Smooth transition for margin change
  },
  drawer: {
    width: 250,
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
    marginTop: '64px', // Adjust for the height of the AppBar
  },
  header: {
    color: '#3f51b5',
  },
  listItem: {
    border: '1px solid #ddd',
    marginBottom: '8px',
  },
  closeButton: {
    margin: '16px',
  },
  drawerClosed: {
    width: "50px",
    marginTop: '64px', // Adjust for the height of the AppBar
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
    alignItems: 'center', // Center the button horizontally
  },
});

function Deployments() {
  const classes = useStyles();
  const [drawerOpen, setDrawerOpen] = useState(true);
  const [deployments, setDeployments] = useState<Deployment[]>([]);

  useEffect(() => {
    const fetchDeployments = async () => {
      try {
        const data = await deploymentsList();
        setDeployments(data);
      } catch (error) {
        console.error('Error fetching deployments:', error);
      }
    };

    fetchDeployments();
  }, []);

  const handleDrawerClose = () => {
    setDrawerOpen(false);
  };

  const handleDrawerOpen = () => {
    setDrawerOpen(true);
  };

  return (
    <>
      <Header />
      <Drawer
        variant="persistent"
        anchor="left"
        open={drawerOpen}
        classes={{ paper: drawerOpen ? classes.drawer : classes.drawerClosed }}
      >
        <List>
          {deployments.map((deployment, index) => (
            <ListItemButton component="li" className={classes.listItem} key={index}>
              <ListItemText primary={deployment.name} />
            </ListItemButton>
          ))}
          {!drawerOpen || (
            <Button variant="contained" color="primary" className={classes.closeButton} onClick={handleDrawerClose}>
              Close Drawer
            </Button>
          )}
        </List>        
      </Drawer>
      {drawerOpen || (<Button variant="contained" color="primary" className={classes.closeButton} onClick={handleDrawerOpen}>
              Open Drawer
            </Button>)}
      <div className={classes.root} style={{ marginLeft: drawerOpen ? '250px' : '50px' }}>
        <h1 className={classes.header}>Details</h1>
      </div>
    </>
  );
}

export default Deployments;