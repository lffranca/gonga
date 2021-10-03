import * as React from 'react';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from "@mui/material/ListSubheader";
import DashboardIcon from '@mui/icons-material/Dashboard';
import InfoIcon from '@mui/icons-material/Info';
import CloudIcon from '@mui/icons-material/Cloud';
import AltRouteIcon from '@mui/icons-material/AltRoute';
import PersonIcon from '@mui/icons-material/Person';
import PowerIcon from '@mui/icons-material/Power';
import ShuffleIcon from '@mui/icons-material/Shuffle';
import SecurityIcon from '@mui/icons-material/Security';
import ListItemLink from "./ListItemLink";

function MenuMainListItems() {
    return (
        <div>
            <ListSubheader inset>API GATEWAY</ListSubheader>
            <ListItemLink to="/" button>
                <ListItemIcon>
                    <DashboardIcon/>
                </ListItemIcon>
                <ListItemText primary="Dashboard"/>
            </ListItemLink>
            <ListItemLink to="/info" button>
                <ListItemIcon>
                    <InfoIcon/>
                </ListItemIcon>
                <ListItemText primary="Information"/>
            </ListItemLink>
            <ListItemLink to="/service" button>
                <ListItemIcon>
                    <CloudIcon/>
                </ListItemIcon>
                <ListItemText primary="Services"/>
            </ListItemLink>
            <ListItemLink to="/route" button>
                <ListItemIcon>
                    <AltRouteIcon/>
                </ListItemIcon>
                <ListItemText primary="Routes"/>
            </ListItemLink>
            <ListItemLink to="/consumer" button>
                <ListItemIcon>
                    <PersonIcon/>
                </ListItemIcon>
                <ListItemText primary="Consumers"/>
            </ListItemLink>
            <ListItemLink to="/plugin" button>
                <ListItemIcon>
                    <PowerIcon/>
                </ListItemIcon>
                <ListItemText primary="Plugins"/>
            </ListItemLink>
            <ListItemLink to="/upstream" button>
                <ListItemIcon>
                    <ShuffleIcon/>
                </ListItemIcon>
                <ListItemText primary="Upstreams"/>
            </ListItemLink>
            <ListItemLink to="/certificate" button>
                <ListItemIcon>
                    <SecurityIcon/>
                </ListItemIcon>
                <ListItemText primary="Certificates"/>
            </ListItemLink>
        </div>
    )
}

export default MenuMainListItems
