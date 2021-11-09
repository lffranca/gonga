import React, {createContext, useState} from "react";

export const GatewayContext = createContext(null);

function GatewayContextProvider(props) {
    const [gateway, setGateway] = useState(true);

    return (
        <GatewayContext.Provider
            value={{
                gateway,
                setGateway
            }}
        >
            {props.children}
        </GatewayContext.Provider>
    );
}

export default GatewayContextProvider
