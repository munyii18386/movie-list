import React from 'react';
import { Route, Switch} from 'react-router-dom';
import {Home} from './containers/Home';
import {NotFound} from "./containers/NotFound";
import {Login} from "./containers/Login";
import {SignUp} from "./containers/SignUp";
import {Logout} from "./containers/Logout";
import {Search} from "./containers/Search";
import {MovieWall} from "./containers/MovieWall"


export const  Routes = () => {
    return(
        <Switch>
            <Route path="/" exact component={Home} />
            <Route path="/login" exact component={Login} />
            <Route path="/signup" exact component={SignUp} />
            <Route path="/logout" exact component={Logout} />
            <Route path="/search" exact component={Search} />
            <Route path="/moviewall" exact component={MovieWall} />
            <Route component={NotFound} />
        </Switch>
        )
}