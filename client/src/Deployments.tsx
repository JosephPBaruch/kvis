import React, { useState } from 'react';
import { Drawer, List, ListItemButton, ListItemText, Button, Box } from "@mui/material";
import { makeStyles } from '@mui/styles';
import Header from './components/Header';

const useStyles = makeStyles({
  root: {
    padding: '20px',
    marginTop: '64px', // Adjust for the height of the AppBar
  },
  drawer: {
    width: 250,
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
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
    width: 50,
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
        classes={{ paper: drawerOpen ? classes.drawer : classes.drawerClosed }}
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
          {drawerOpen ? (
            <Button variant="contained" color="primary" className={classes.closeButton} onClick={handleDrawerClose}>
              Close Drawer
            </Button>
          ) : (
            <Button variant="contained" color="primary" className={classes.closeButton} onClick={handleDrawerOpen}>
              Open Drawer
            </Button>
          )}
        </Box>
      </Drawer>
      <div className={classes.root}>
        <h1 className={classes.header}>Details</h1>
      </div>
    </>
  );
}

export default Deployments;