import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import TableRow from './components/TableRow';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {serverports: []};
  }
  componentDidMount(){
    axios.get('http://localhost:9999/api')
    .then(response => {
      this.setState({ serverports: response.data });
    })
    .catch(function (error) {
      console.log(error);
    })
  }
  tabRow(){
      return this.state.serverports.map(function(object, i){
          return <TableRow obj={object} key={i} />;
      });
  }

  render() {
    return (
      <div className="container">
          <table className="table table-striped">
            <thead>
              <tr>
                <td>ID</td>
                <td>Name</td>
                <td>Port</td>
              </tr>
            </thead>
            <tbody>
              {this.tabRow()}
            </tbody>
          </table>
      </div>
    );
  }
}

export default App;
