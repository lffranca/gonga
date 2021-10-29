import * as React from "react";
import {useContext} from "react";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import Divider from "@mui/material/Divider";
import {List} from "@mui/material";
import MenuMainListItems from "./MenuMainListItems";
import MenuAdminListItems from "./MenuAdminListItems";
import {styled} from "@mui/material/styles";
import MuiDrawer from "@mui/material/Drawer";
import {SidebarContext} from "../state/SidebarContext";

const drawerWidth = 240;


const Drawer = styled(MuiDrawer, {shouldForwardProp: (prop) => prop !== 'open'})(
    ({theme, open}) => ({
        '& .MuiDrawer-paper': {
            position: 'relative',
            whiteSpace: 'nowrap',
            height: '100vh',
            width: drawerWidth,
            transition: theme.transitions.create('width', {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen,
            }),
            boxSizing: 'border-box',
            ...(!open && {
                overflowX: 'hidden',
                transition: theme.transitions.create('width', {
                    easing: theme.transitions.easing.sharp,
                    duration: theme.transitions.duration.leavingScreen,
                }),
                width: theme.spacing(7),
                [theme.breakpoints.up('sm')]: {
                    width: theme.spacing(9),
                },
            }),
        },
    }),
);

function Sidebar() {
    const {open, setOpen} = useContext(SidebarContext)
    const toggleDrawer = () => {
        setOpen(!open);
    };

    return (
        <Drawer variant="permanent" open={open}>
            <Toolbar
                sx={{
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'flex-end',
                    px: [1],
                }}
            >
                <IconButton onClick={toggleDrawer}>
                    <ChevronLeftIcon/>
                </IconButton>
            </Toolbar>
            <Divider/>
            <div style={{height: "calc(100% - 64px)", overflow: 'auto'}}>
                <List>
                    <MenuMainListItems/>
                </List>
                <Divider/>
                <List>
                    <MenuAdminListItems/>
                </List>
            </div>
        </Drawer>
    )
}

export default Sidebar