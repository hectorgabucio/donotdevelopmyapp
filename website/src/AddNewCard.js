import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import CircularProgress from '@material-ui/core/CircularProgress';
import IconButton from '@material-ui/core/IconButton';
import CasinoIcon from '@material-ui/icons/Casino';

const useStyles = makeStyles((theme) => ({
  root: {
    width: 345,
    height: 450,
    margin: 5,
  },
  notFoundRoot: {
    display: 'flex',
    width: 400,
  },
  details: {
    display: 'flex',
    flexDirection: 'column',
  },
  content: {
    flex: '1 0 auto',
  },
  cover: {
    width: 350,
  },
  controls: {
    display: 'flex',
    alignItems: 'center',
    paddingLeft: theme.spacing(1),
    paddingBottom: theme.spacing(1),
  },
  retryIcon: {
    height: 38,
    width: 38,
  },
}));

export const AddNewCard = () => {
  const classes = useStyles();
  const [character, setCharacter] = useState(null);
  const [loaded, setLoaded] = useState(false);
  useEffect(() => {
    getNewCharacter();
  }, []);

  async function getNewCharacter() {
    try {
      const character = await addNewCharacter();
      setCharacter(character);
    } finally {
      setLoaded(true);
    }
  }

  if (!loaded) {
    return <CircularProgress color="secondary" />;
  }

  if (!character) {
    return (
      <Card className={classes.notFoundRoot}>
        <div className={classes.details}>
          <CardContent className={classes.content}>
            <Typography component="h5" variant="h5">
              Unlucky! no character found, try again.
            </Typography>
          </CardContent>
          <div className={classes.controls}>
            <IconButton aria-label="retry" onClick={getNewCharacter}>
              <CasinoIcon className={classes.retryIcon} />
            </IconButton>
          </div>
        </div>
        <CardMedia
          className={classes.cover}
          image="/sad-morty.webp"
          title="Sad Morty"
        />
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
