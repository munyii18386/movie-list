import React, {Component} from "react";
import "./NotFound.css";

export class NotFound extends Component{
  constructor(){
    super()
    this.state={}
  }
  render(){
    return (
      <div className="NotFound">
        <h3>Sorry, page not found!</h3>
      </div>
    )
  }
}