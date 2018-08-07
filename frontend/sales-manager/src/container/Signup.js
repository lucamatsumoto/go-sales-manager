import React, { Component } from 'react'
import { FormGroup, FormControl, ControlLabel, HelpBlock } from 'react-bootstrap'
import LoaderButton from '../components/LoaderButton'
import "./Signup.css"

export default class Signup extends Component {
    constructor(props) {
        super(props);
        
        this.state = {
            name: "",
            email: "",
            password: "",
            confirmPassword: "",
            newUser: null,
            isLoading: false,
        };
    }

    validateForm() {
        return (
            this.state.name.length >= 0 &&
            this.state.name.length >= 0 &&
            this.state.password.length >= 0 &&
            this.state.password == this.state.confirmPassword
        )
    }

    handleSubmit = async event => {
        event.preventDefault()
        this.setState({ isLoading: true });
        this.setState({ isLoading: false });    
    }

    handleChange = event => {
        this.setState({
            [event.target.id]: event.target.value
        })
    }

    renderForm() {
        return (
            <form onSubmit={this.handleSubmit}>
                <FormGroup controlId="name" beSize="large">
                <ControlLabel>Name</ControlLabel>
                <FormControl
                    autoFocus
                    value={this.state.name}
                    onChange={this.handleChange}
                />
                </FormGroup>
                <FormGroup controlId="email" beSize="large">
                <ControlLabel>Email</ControlLabel>
                <FormControl
                    autoFocus
                    type="email"
                    value={this.state.email}
                    onChange={this.handleChange}
                />
                </FormGroup>
                <FormGroup controlId="password" beSize="large">
                <ControlLabel>Password</ControlLabel>
                <FormControl
                    autoFocus
                    type="password"
                    value={this.state.password}
                    onChange={this.handleChange}
                />
                </FormGroup>
                <FormGroup controlId="confirmPassword" beSize="large">
                <ControlLabel>ConfirmPassword</ControlLabel>
                <FormControl
                    autoFocus
                    type="password"
                    value={this.state.confirmPassword}
                    onChange={this.handleChange}
                />
                </FormGroup>
                <LoaderButton
                    block
                    bsSize="large"
                    disabled={!this.validateForm()}
                    type="submit"
                    isLoading={this.state.isLoading}
                    text="Signup"
                    loadingText="Signing up…"
                />
            </form> 
        )
    }

    render() {
        return (
          <div className="Signup">
            {this.renderForm()}
          </div>
        );
      }
}
