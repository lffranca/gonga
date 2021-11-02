import React, {forwardRef, useMemo} from "react";
import {Link as RouterLink} from "react-router-dom";
import {Link} from "@mui/material";

const ItemLink = ({to, ...propsToItem}) => {
    const renderLink = useMemo(
        () =>
            forwardRef((itemProps, ref) => (
                <RouterLink to={to} ref={ref} {...itemProps} />
            )),
        [to]
    );

    return (
        <Link component={renderLink} {...propsToItem}>
            {propsToItem.children}
        </Link>
    );
};

export default ItemLink
