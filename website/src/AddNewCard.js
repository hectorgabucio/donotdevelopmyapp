import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import CircularProgress from '@material-ui/core/CircularProgress';

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
});

export const AddNewCard = () => {
  const classes = useStyles();
  const [character, setCharacter] = useState(null);
  const [loaded, setLoaded] = useState(false);
  useEffect(() => {
    async function getNewCharacter() {
      try {
        const character = await addNewCharacter();
        setCharacter(character);
      } finally {
        setLoaded(true);
      }
    }
    getNewCharacter();
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
