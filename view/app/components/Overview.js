import * as React from "react";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";

function Overview() {
    return (
        <Grid container spacing={3}>
            {/* Chart */}
            <Grid item xs={12} md={8} lg={9}>
                <Paper
                    sx={{
                        p: 2,
                        display: 'flex',
                        flexDirection: 'column',
                        height: 240,
                    }}
                >
                    {/*<Chart/>*/}
                    {"Chart"}
                </Paper>
            </Grid>
            {/* Recent Deposits */}
            <Grid item xs={12} md={4} lg={3}>
                <Paper
                    sx={{
                        p: 2,
                        display: 'flex',
                        flexDirection: 'column',
                        height: 240,
                    }}
                >
                    {/*<Deposits/>*/}
                    {"Deposits"}
                </Paper>
            </Grid>
            {/* Recent Orders */}
            <Grid item xs={12}>
                <Paper sx={{p: 2, display: 'flex', flexDirection: 'column'}}>
                    {/*<Orders/>*/}
                    {"Orders"}
                </Paper>
            </Grid>
        </Grid>
    )
}

export default Overview
