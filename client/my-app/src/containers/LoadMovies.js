import React  from "react";
import {Container, Nav, Row, Col, Form, Button, Card, CardGroup} from "react-bootstrap";
import axios from "axios";
import "./MovieWall.css";
import { useSelector } from "react-redux";
import { useDispatch } from "react-redux";

 export const LoadMovies = () => {

    // console.log("something")
    const state  = useSelector(state => state);
    const dispatch = useDispatch();

    const initialState = {
        results: [],
        updated: false
    };
    const [data, setData] = React.useState(initialState);

    let config = {
        headers: {
            Authorization: state.token
        }
    }
    axios.get('https://localhost/api/GetWall', config)
    .then((response)=>{
        // console.log(response)
        // console.log(response.data)
        // console.log(data.results)
        setData({
            ...data,
            results: response.data,
            updated: true
        })
        


    }) 
    .catch((e) =>{
        console.log(e)
    })

   
    const list = data.results.map((item, index) => 
                <div key={index} className="search-margin">
                    <Col md={4} >
                        <Card key={index} style={{ width: '18rem' }}>
                            <Card.Img variant="top" src={item.movie_url}/>
                            <Card.Body>
                                <Card.Title>{item.title}</Card.Title>
                                <Card.Text>
                                {item.overview}
                                </Card.Text>
                            </Card.Body>
                        </Card>
                     </Col>
                </div> 
        
    )

    
    const empty = 
        <div  className="search-margin">
            <Col md={4} >
                <h1>Wall Under Construction...</h1>
            </Col>
        </div> 
    
        
   


   
        return(
           data.results == undefined || data.results == null ? empty : list
        // <h1>Something</h1>
        )
   
   
   
 
          
   

}