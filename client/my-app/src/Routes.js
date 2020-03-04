import React from 'react';
import { Route, Switch} from 'react-router-dom';
import {Home} from './containers/Home';
import {NotFound} from "./containers/NotFound";
import {Login} from "./containers/Login";
import {SignUp} from "./containers/SignUp";
import {Logout} from "./containers/Logout";
import {MyWall} from "./containers/MyWall";
import {Community} from "./containers/Community"
import {AppContext} from './App';

export const  Routes = () => {
    const {state , dispatch } = React.useContext(AppContext);
    return(
        <Switch>
            <Route path="/" exact component={Home} />
            <Route path="/login" exact component={Login} />
            <Route path="/signup" exact component={SignUp} />
            <Route path="/logout" exact component={Logout} />
            <Route path="/mywall" exact component={MyWall} />
            <Route component={NotFound} />
        </Switch>
        )
}