import React from 'react'
import {createTheme, ThemeProvider} from "@mui/material";
import AuthContextProvider from "./AuthContext";
import SidebarContextProvider from "./SidebarContext";
import GatewayContextProvider from "./GatewayContext";

const mdTheme = createTheme();

function State(props) {
    return (
        <ThemeProvider theme={mdTheme}>
            <AuthContextProvider>
                <SidebarContextProvider>
                    <GatewayContextProvider>
                        {props.children}
                    </GatewayContextProvider>
                </SidebarContextProvider>
            </AuthContextProvider>
        </ThemeProvider>
    )
}

export default State