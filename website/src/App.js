import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import axios from 'axios';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

export default function App() {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/about">About</Link>
            </li>
            <li>
              <Link to="/users">Users</Link>
            </li>
          </ul>
        </nav>

        {/* A <Switch> looks through its children <Route>s and
            renders the first one that matches the current URL. */}
        <Switch>
          <Route path="/about">
            <About />
          </Route>
          <Route path="/users">
            <Users />
          </Route>
          <Route path="/">
            <Dashboard />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

/*
function Home() {
  axios.get('/random').then(
    (onfullfilled) => {
      console.log(onfullfilled);
    },
    (rejected) => {
      console.log(rejected);
    }
  );

  return (
    <Button variant="contained" color="primary">
      Hello World
    </Button>
  );
}

*/

const Dashboard = () => {
  const classes = useStyles();
  const [character, setCharacter] = useState(null);
  useEffect(() => {
    async function getCharacter() {
      const character = await getRandomCharacter();
      setCharacter(character);
    }
    getCharacter();
  }, []);

  if (!character) {
    return (
      <Button variant="contained" color="primary">
        Hello World
      </Button>
    );
  } else {
    return (
      <Card className={classes.root}>
        <CardActionArea>
          <CardMedia
            component="img"
            alt="Ricky and Morty image"
            image={character.image}
            title={character.name}
          />
          <CardContent>
            <Typography gutterBottom variant="h5" component="h2">
              {character.name}
            </Typography>
          </CardContent>
        </CardActionArea>
        <CardActions>
          <Button size="small" color="primary">
            Share
          </Button>
          <Button size="small" color="primary">
            Learn More
          </Button>
        </CardActions>
      </Card>
    );
  }
};

async function getRandomCharacter() {
  const resp = await axios.get('/random');
  return resp.status === 200 ? resp.data : null;
}

function About() {
  return <h2>About</h2>;
}

function Users() {
  return <h2>Users</h2>;
}
