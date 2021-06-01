import { Component } from 'react';
import { Navbar, Nav, Container } from 'react-bootstrap'

export default class Main extends Component {
    render() {
        return (
            <div>
                <div>
                    <Navbar bg="dark" variant="dark" expand="lg">
                      <Navbar.Brand href="/">Gitploy</Navbar.Brand>
                      <Navbar.Toggle aria-controls="basic-navbar-nav" />
                      <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="mr-auto">
                          <Nav.Link href="/">Home</Nav.Link>
                          <Nav.Link href="/history">History</Nav.Link>
                        </Nav>
                      </Navbar.Collapse>
                    </Navbar>
                </div>
                <Container className="mt-5">
                    {this.props.children}
                </Container>
            </div>
        )
    }
}