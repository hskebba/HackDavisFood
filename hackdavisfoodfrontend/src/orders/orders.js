import React from "react"
import { Button, CloseButton, Container, Row, Col, Card } from 'react-bootstrap';

class Orders extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            data: [
                {
                    "id": "1",
                    "description": "Generic Description",
                    "pickup_at": "5:00PM",
                    "location": "My House",
                    "supplier": "Hayley",
                    "status": "Open",
                    "image": "https://www.rd.com/wp-content/uploads/2021/03/GettyImages-1133605325-scaled-e1617227898456.jpg",
                    "updated_at": "Now",
                    "created_at": "Now2"
                },
                {
                    "id": "2",
                    "description": "Generic Description 2",
                    "pickup_at": "5:00PM",
                    "location": "In the Middle of The Street",
                    "supplier": "Ian",
                    "status": "Closed",
                    "image": "https://cdn.mos.cms.futurecdn.net/KYEJp9vem3QQFGhi25SYx4-1200-80.jpg",
                    "updated_at": "Now",
                    "created_at": "Now2"
                }
            ]
        };
    }

    render() {
        return (
            <Container className="mt-3">
                <h1>
                    Orders
                </h1>
                <Row xs={1} md={2} lg={3} className="g-4">
                    {Array.from({ length: this.state.data.length }).map((_, idx) => (
                        <Col className="d-flex align-items-stretch">
                            <Card>
                                <CloseButton style={{position:"absolute",right:"0px"}}/>
                                <Card.Img variant="top" src={this.state.data[idx].image} />
                                <Card.Body>
                                    <Card.Title>{this.state.data[idx].supplier} - {this.state.data[idx].location} - {this.state.data[idx].pickup_at}</Card.Title>
                                    <Card.Text>{this.state.data[idx].description}</Card.Text>
                                </Card.Body>
                                <Container className='mb-3'>
                                <Row>
                                    <Col>
                                        <Button variant="primary" size="sm" active>
                                            Claim
                                        </Button>{' '}
                                    </Col>
                                    <Col>
                                        <Button variant="success" size="sm" active>
                                            Complete
                                        </Button>
                                    </Col>
                                </Row>
                                </Container>
                            </Card>
                        </Col>
                    ))}
                </Row>
            </Container>
        )
    }
}

export {Orders}