import React, {createContext, useState} from "react";

export const AuthContext = createContext(null);

function AuthContextProvider(props) {
    const [authenticated, setAuthenticated] = useState(true);

    return (
        <AuthContext.Provider
            value={{
                authenticated,
                setAuthenticated
            }}
        >
            {props.children}
        </AuthContext.Provider>
    );
}

export default AuthContextProvider
