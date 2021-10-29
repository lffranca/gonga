import * as React from 'react';
import {BrowserRouter as Router} from "react-router-dom";
import Box from '@mui/material/Box';
import Header from "./Header";
import Sidebar from "./Sidebar";
import Routes from "./Routes";

function Private() {
    return (
        <Router basename="/app">
            <Box sx={{display: 'flex'}}>
                <Header/>
                <Sidebar/>
                <Routes/>
            </Box>
        </Router>
    );
}

export default Private