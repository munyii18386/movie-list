import React, {Component} from 'react';
import './Home.css';

export class  Home extends Component{
    constructor(){
        super()
        this.state={}
    }
    render(){
        return(
            <div className="Home">
                <div className="landing">
                    <h1>Wall Of Movies</h1>
                    <p>A simple movie list app</p>
                </div>
            </div>
        )
    }
}