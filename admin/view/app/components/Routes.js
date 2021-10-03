import React from 'react'
import {Route, Switch} from "react-router-dom";
import DashboardPage from "../pages/DasboardPage";
import InfoPage from "../pages/InfoPage";
import ServicePage from "../pages/ServicePage";
import RoutePage from "../pages/RoutePage";
import ConsumerPage from "../pages/ConsumerPage";
import PluginPage from "../pages/PluginPage";
import UpstreamPage from "../pages/UpstreamPage";
import CertificatePage from "../pages/CertificatePage";
import UserPage from "../pages/UserPage";
import ConnectionPage from "../pages/ConnectionPage";
import SettingsPage from "../pages/SettingsPage";
import SnapshotPage from "../pages/SnapshotPage";

function Routes() {
    return (
        <Switch>
            <Route exact path="/">
                <DashboardPage/>
            </Route>
            <Route path="/info">
                <InfoPage/>
            </Route>
            <Route path="/service">
                <ServicePage/>
            </Route>
            <Route path="/route">
                <RoutePage/>
            </Route>
            <Route path="/consumer">
                <ConsumerPage/>
            </Route>
            <Route path="/plugin">
                <PluginPage/>
            </Route>
            <Route path="/upstream">
                <UpstreamPage/>
            </Route>
            <Route path="/certificate">
                <CertificatePage/>
            </Route>

            <Route path="/user">
                <UserPage/>
            </Route>
            <Route path="/connection">
                <ConnectionPage/>
            </Route>
            <Route path="/snapshot">
                <SnapshotPage/>
            </Route>
            <Route path="/settings">
                <SettingsPage/>
            </Route>
        </Switch>
    )
}

export default Routes
