import React from 'react';
import Card from '@material-ui/core/Card';
import CasinoIcon from '@material-ui/icons/Casino';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import sadMorty from '../images/sad-morty.webp';
import { PropTypes } from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  limitReachedRoot: {
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

const LimitReached = (props) => {
  const classes = useStyles();
  return (
    <Card className={classes.limitReachedRoot}>
      <div className={classes.details}>
        <CardContent className={classes.content}>
          <Typography component="h5" variant="h5">
            Limit reached! You cant get new cards now, wait until next day to
            get new cards.
          </Typography>
        </CardContent>
      </div>
      <CardMedia className={classes.cover} image={sadMorty} title="Sad Morty" />
    </Card>
  );
};

export default LimitReached;
