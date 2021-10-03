import React, {useContext} from 'react'
import {AuthContext} from "../state/AuthContext";
import Private from "./Private";
import Public from "./Public";

function Base() {
    const {authenticated} = useContext(AuthContext)
    return authenticated ? <Private/> : <Public/>
}

export default Base
