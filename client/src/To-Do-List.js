import React,{Component} from "React";
import axios from "axios";  //used for the api calls
import {Card, Header, Form , Input,Icon} from "semantic-ui-react";

let endpoint = "http://localhost:9000";

class ToDoList extends Component{
    constructor(props){
        super(props);
        this.state={
            task:"",
            items:[],
        };

    }
    ComponentDidMount(){
        this.getTask();
    }
    
    rander()
}



