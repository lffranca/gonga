import React from 'react'
import Base from "./components/Base";
import CssBaseline from "@mui/material/CssBaseline";
import State from "./state/State";

function App() {
    return (
        <State>
            <CssBaseline/>
            <Base/>
        </State>
    );
}

export default App;
