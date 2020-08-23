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
import CircularProgress from '@material-ui/core/CircularProgress';

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
  notFoundRoot: {
    minWidth: 345,
  },
});

export const Dashboard = () => {
  const classes = useStyles();
  const [character, setCharacter] = useState(null);
  const [loaded, setLoaded] = useState(false);
  useEffect(() => {
    async function getCharacter() {
      try {
        const character = await getRandomCharacter();
        setCharacter(character);
      } finally {
        setLoaded(true);
      }
    }
    getCharacter();
  }, []);

  if (!loaded) {
    return <CircularProgress color="secondary" />;
  }

  if (!character) {
    return (
      <Card className={classes.notFoundRoot}>
        <CardContent>
          <Typography variant="h5" component="h2">
            Unlucky! no character found, try again refreshing.
          </Typography>
        </CardContent>
      </Card>
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
  const resp = await axios.post('/characters');
  return resp.status === 200 ? resp.data : null;
}
