import * as React from 'react';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import ListItemIcon from '@mui/material/ListItemIcon';
import PeopleIcon from '@mui/icons-material/People';
import CastIcon from '@mui/icons-material/Cast';
import PhotoCameraIcon from '@mui/icons-material/PhotoCamera';
import SettingsIcon from '@mui/icons-material/Settings';
import ListItemLink from "./ListItemLink";

function MenuAdminListItems() {
    return (
        <div>
            <ListSubheader inset>APPLICATION</ListSubheader>
            <ListItemLink to="/user" button>
                <ListItemIcon>
                    <PeopleIcon/>
                </ListItemIcon>
                <ListItemText primary="Users"/>
            </ListItemLink>
            <ListItemLink to="/connection" button>
                <ListItemIcon>
                    <CastIcon/>
                </ListItemIcon>
                <ListItemText primary="Connections"/>
            </ListItemLink>
            <ListItemLink to="/snapshot" button>
                <ListItemIcon>
                    <PhotoCameraIcon/>
                </ListItemIcon>
                <ListItemText primary="Snapshots"/>
            </ListItemLink>
            <ListItemLink to="settings" button>
                <ListItemIcon>
                    <SettingsIcon/>
                </ListItemIcon>
                <ListItemText primary="Settings"/>
            </ListItemLink>
        </div>
    )
}

export default MenuAdminListItems