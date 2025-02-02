import { useState, useEffect } from 'react';
import { Drawer, List, ListItemButton, ListItemText, Button } from "@mui/material";
import { makeStyles } from '@mui/styles';
import Header from './components/Header';
import { ListObjects } from './utils/http';
import { ListObject, KubePageProps } from './types/types';
import DetailInformation from './components/DetailInfo';

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
  selectedItem: {
    backgroundColor: '#e0e0e0', // Change background color to indicate selection
  },
});

const KubePage: React.FC<KubePageProps> = ({ typeOption }) => {
  const classes = useStyles();
  const [drawerOpen, setDrawerOpen] = useState(true);
  const [list, setList] = useState<ListObject[]>([]);
  const [selected, setSelected] = useState<ListObject | null>(null);

  useEffect(() => {
    const fetchList = async () => {
      try {
        const data: ListObject[] = await ListObjects(typeOption);
        setList(data);
      } catch (error) {
        console.error('Error fetching list:', error);
      }
    };

    fetchList();
  }, [typeOption]);

  const handleDrawerClose = () => {
    setDrawerOpen(false);
  };

  const handleDrawerOpen = () => {
    setDrawerOpen(true);
  };

  const handleListItemClick = (item: ListObject) => {
    setSelected(item);
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
          {list && list.map((item, index) => (
            <ListItemButton
              component="li"
              className={`${classes.listItem} ${selected === item ? classes.selectedItem : ''}`}
              key={index}
              onClick={() => handleListItemClick(item)}
            >
              <ListItemText primary={item.name} />
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
        {selected && (
          <DetailInformation typeOption={typeOption} name={selected.name} />
        )}
      </div>
    </>
  );
}

export default KubePage;