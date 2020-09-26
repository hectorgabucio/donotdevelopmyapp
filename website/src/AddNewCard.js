import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import CircularProgress from '@material-ui/core/CircularProgress';

import NotFoundCard from './components/NotFoundCard';
import LimitReached from './components/LimitReached';

const useStyles = makeStyles((theme) => ({
  root: {
    width: 345,
    height: 450,
    marginLeft: 30,
  },
}));

export const AddNewCard = () => {
  const classes = useStyles();
  const [character, setCharacter] = useState(null);
  const [limitReached, setLimitReached] = useState(false);
  const [loaded, setLoaded] = useState(false);
  useEffect(() => {
    getNewCharacter();
  }, []);

  async function getNewCharacter() {
    try {
      const character = await addNewCharacter();
      setCharacter(character);
    } catch (error) {
      // Error 😨
      if (error.response) {
        console.log('server error');
        if (error.response.status === 503) {
          console.log('limit reached');
          setLimitReached(true);
        }
      } else if (error.request) {
        /*
         * The request was made but no response was received, `error.request`
         * is an instance of XMLHttpRequest in the browser and an instance
         * of http.ClientRequest in Node.js
         */
        console.log('client error');
        console.log(error.request);
      } else {
        // Something happened in setting up the request and triggered an Error
        console.log('something happened');
        console.log('Error', error.message);
      }
      console.log(error);
    } finally {
      setLoaded(true);
    }
  }

  if (!loaded) {
    return <CircularProgress color="secondary" />;
  }

  if (limitReached) {
    return <LimitReached></LimitReached>;
  } else if (!character) {
    return <NotFoundCard retry={getNewCharacter}></NotFoundCard>;
  } else {
    return (
      <Card key={character.toString()} className={classes.root}>
        <CardActionArea>
          <CardMedia
            component="img"
            alt="Ricky and Morty image"
            image={character.image}
            title={character.name}
          />
          <CardContent>
            <Typography gutterBottom component="h3">
              {character.name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    );
  }
};

async function addNewCharacter() {
  const resp = await axios.post('/characters');
  return resp.status === 200 ? resp.data : null;
}
