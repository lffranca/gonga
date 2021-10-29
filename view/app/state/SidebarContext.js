import React, {createContext, useState} from "react";

export const SidebarContext = createContext(null);

function SidebarContextProvider(props) {
    const [open, setOpen] = useState(true);

    return (
        <SidebarContext.Provider
            value={{
                open,
                setOpen
            }}
        >
            {props.children}
        </SidebarContext.Provider>
    );
}

export default SidebarContextProvider
