import React from 'react';
import { Route, Switch} from 'react-router-dom';
import Home from './containers/Home';
import NotFound from "./containers/NotFound";
import {Login} from "./containers/Login";
import SignUp from "./containers/SignUp";
import MyWall from "./containers/MyWall";
import Community from "./containers/Community"

export default function Routes() {
    return(
        <Switch>
            <Route path="/" exact component={Home} />
            <Route path="/login" exact component={Login} />
            <Route path="/signup" exact component={SignUp} />
            <Route component={NotFound} />
        </Switch>
    )
}