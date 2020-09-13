import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';
import CircularProgress from '@material-ui/core/CircularProgress';
import CharacterCard from './components/CharacterCard';

const useStyles = makeStyles({
  root: {
    width: 200,
    height: 300,
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
    justifyContent: 'center',
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
            :( You dont have any cards! Pick new cards in the left menu.
          </Typography>
        </CardContent>
      </Card>
    );
  } else {
    return (
      <div className={classes.cardsContainer}>
        {characters.map((character) => {
          return (
            <CharacterCard
              key={character.toString()}
              image={character.character.image}
              name={character.character.name}
            ></CharacterCard>
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
