import React from "react";
import axios from "axios";
import {Container} from "react-bootstrap";
import { Redirect } from 'react-router-dom'
import { useSelector, useDispatch} from "react-redux";





export const Logout = () => {

    const state = useSelector(state => state)
    const dispatch = useDispatch()

    console.log("SignUp user is auth: " + state.isAuthenticated)

    axios({
        method: 'patch', //you can set what request you want to be
        url: 'https://localhost/api/Logout',
        headers: {
          Authorization: state.token
        }
    })
    .then((response) =>{
        console.log(response)
        dispatch({
            type: "LOGOUT"
        })
    })
    .catch((e) =>{
        console.log(e)
    })
        return(
            <Container>
                <Redirect to='/' />
            </Container>
        )
   
}



export default Logout