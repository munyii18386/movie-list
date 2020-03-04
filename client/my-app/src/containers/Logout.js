import React  from "react";
import {AppContext} from '../App';
import {Container,Button, Row, Col, ButtonGroup} from "react-bootstrap";

export const Logout = () => {

    return(
        <Container>
            <Row>
                <p>Are you sure you want to logout?</p>
            </Row>
            <Row>
                <Col>
                <ButtonGroup aria-label="Basic example">
                    <Button variant="secondary">YES</Button>
                    <Button disabled="true" variant="secondary"></Button>
                    <Button variant="secondary">NO</Button>
                </ButtonGroup>
                </Col>
            </Row>
        </Container>
    )

}

export default Logout