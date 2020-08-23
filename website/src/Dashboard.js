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
import { FormHelperText } from '@material-ui/core';

const useStyles = makeStyles({
  root: {
    width: 345,
    height: 450,
    margin: 5,
  },
  notFoundRoot: {
    width: 345,
    height: 450,
    margin: 5,
  },
  cardsContainer: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'space-evenly',
    alignItems: 'center',
  },
});

export const Dashboard = () => {
  const classes = useStyles();
  const [characters, setCharacters] = useState(null);
  const [loaded, setLoaded] = useState(false);
  useEffect(() => {
    async function getCharacters() {
      try {
        const characters = await getMyCharacters();
        setCharacters(characters);
      } finally {
        setLoaded(true);
      }
    }
    getCharacters();
  }, []);

  if (!loaded) {
    return <CircularProgress color="secondary" />;
  }

  if (!characters) {
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
    console.log(characters);

    return (
      <div className={classes.cardsContainer}>
        {characters.map((character) => {
          return (
            <Card key={character.toString()} className={classes.root}>
              <CardActionArea>
                <CardMedia
                  component="img"
                  alt="Ricky and Morty image"
                  image={character.character.image}
                  title={character.character.name}
                />
                <CardContent>
                  <Typography gutterBottom variant="h5" component="h2">
                    {character.character.name}
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
        })}
      </div>
    );
  }
};

async function addNewCharacter() {
  const resp = await axios.post('/characters');
  return resp.status === 200 ? resp.data : null;
}

async function getMyCharacters() {
  const resp = await axios.get('/characters/me');
  return resp.status === 200 ? resp.data : null;
}
