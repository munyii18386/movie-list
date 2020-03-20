import React  from "react";
import {Container, Nav, Row, Col, Form, Button, Card, CardGroup} from "react-bootstrap";
import axios from "axios";
import "./Search.css";
import { useSelector } from "react-redux";
import { useDispatch } from "react-redux";

export const  Search = () => {
    const state  = useSelector(state => state);
    const dispatch = useDispatch();

    const initialState = {
        keyphrase: "",
        results: [],
        cardID: -1,
        index: -1,
        text: "Add to Wall",
    };
    const [data, setData] = React.useState(initialState);

    const handleInputChange = event => {
        setData({
            ...data,
            [event.target.name]: event.target.value
          });
    }
    
    
    const handleSubmit = event => {
        event.preventDefault()
        const url = 'https://api.themoviedb.org/3/search/movie?api_key=b80cde793a3f70336b85099500977b28&language=en-US&query=' + data.keyphrase + '&page=1&include_adult=false'
        console.log(url)
        axios.get(url)
        .then((response)=>{
            // console.log(response)
            // console.log(response.data.results)
            setData({
                ...data,
                results: response.data.results
            })

        }) 
        .catch((e) =>{
            console.log(e)
        })  
        
    }

    const handleAdd = event => {
        event.preventDefault()
        console.log(event.target.id)
        const item = data.results[event.target.id]
        console.log(event.target.textContent)
        data.cardID = event.target.id
        
        // console.log(item)
        let url = ""
         if(item.poster_path ==  null){
            url = require('../img/reel.jpg')
        } else{
            url = 'https://image.tmdb.org/t/p/w500/'+item.poster_path
        }
        const movie = {
            id: event.target.id,
            url: url,
            title: item.title, 
            overview: item.overview
        }



        let config = {
            headers: {
                Authorization: state.token
            }
        }
        
        axios.post('https://localhost/api/Wall', movie, config)
        .then((response)=>{
            console.log(response)
            console.log(response.data)
            console.log(data)
            if(response.data.movie_added){
                let  btn = document.getElementById(data.cardID);
                btn.textContent = "Added"
            }
          
            

        }) 
        .catch((e) =>{
            console.log(e)
            alert("Unable to add movie")
        }) 

       
    }

    // if (data.cardList.length > 0){
    //     data.cardList.map(item => {
    //         // 
    //         console.log(item)
    //         console.log(typeof data.cardID.toString())
    //         console.log(item.localeCompare(data.cardID.toString()))
    //         if(item.localeCompare(data.cardID)){
    //                setData({
    //                    text: "Added"
    //                })
    //         }
    //     })
    // }

        const list = data.results.map((item, index) => {
            if(data.results && data.results.length){
                data.index = index

                
                let url = ""
                if(item.poster_path ==  null){
                    url = require('../img/reel.jpg')
                } else{
                    url = 'https://image.tmdb.org/t/p/w500'+item.poster_path
                }
                // console.log("find the title by id: " + data.results[index].title)

                return(
                    <div className="search-margin">
                        <Col md={4} >
                            <Card style={{ width: '18rem' }}>
                                <Card.Img variant="top" src={url}/>
                                <Card.Body>
                                    <Card.Title>{item.title}</Card.Title>
                                    <Card.Text>
                                    {item.overview}
                                    </Card.Text>
                                    <Button  id={data.index} onClick={handleAdd} variant="primary">{data.text}</Button>
                                </Card.Body>
                            </Card>
                        </Col>
                    </div> 
                    
                )

            }
        })

        console.log(data)

   

        return(
            <Container className="search-container">
                <div className="search-margin"></div>
                <Nav fill variant="tabs" defaultActiveKey="/search">
                    <Nav.Item>
                        <Nav.Link href="/search">Search</Nav.Link>
                    </Nav.Item>
                    <Nav.Item>
                        <Nav.Link  href="/moviewall">Movie Wall</Nav.Link>
                    </Nav.Item>
               
                </Nav>
                <Form onSubmit={handleSubmit} id="search-bar" >
                <Form.Row>
                    <Col>
                        <Form.Control
                            value={data.keyphrase}
                            onChange={handleInputChange}
                            name="keyphrase"
                            placeholder="Search Titles, People, Genres.." 
                        />
                    </Col>
                    <Button variant="primary" type="submit">Search</Button>
                </Form.Row>
            </Form>
            <div className="search-margin"></div>
            
                <Row>
                    {list}
                </Row>
           
            </Container>
           
        )

}