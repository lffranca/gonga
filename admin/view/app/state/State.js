import React from 'react'
import {createTheme, ThemeProvider} from "@mui/material";
import AuthContextProvider from "./AuthContext";
import SidebarContextProvider from "./SidebarContext";

const mdTheme = createTheme();

function State(props) {
    return (
        <ThemeProvider theme={mdTheme}>
            <AuthContextProvider>
                <SidebarContextProvider>
                    {props.children}
                </SidebarContextProvider>
            </AuthContextProvider>
        </ThemeProvider>
    )
}

export default State