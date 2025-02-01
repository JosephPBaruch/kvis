import React, { useState } from 'react';
import { Drawer, List, ListItemButton, ListItemText, Button, Box } from "@mui/material";
import { makeStyles } from '@mui/styles';
import Header from './components/Header';

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
        className={drawerOpen ? classes.drawer : classes.drawerClosed }
      >
        <List>
          <ListItemButton component="li" className={classes.listItem}>
            <ListItemText primary="Item 1" />
          </ListItemButton>
          <ListItemButton component="li" className={classes.listItem}>
            <ListItemText primary="Item 2" />
          </ListItemButton>
          <ListItemButton component="li" className={classes.listItem}>
            <ListItemText primary="Item 3" />
          </ListItemButton>
        </List>
        <Box>
          {!drawerOpen || (
            <Button variant="contained" color="primary" className={classes.closeButton} onClick={handleDrawerClose}>
              Close Drawer
            </Button>
          )}
        </Box>
        
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