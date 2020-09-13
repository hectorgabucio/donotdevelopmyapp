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

const NotFoundCard = (props) => {
  const classes = useStyles();
  return (
    <Card className={classes.notFoundRoot}>
      <div className={classes.details}>
        <CardContent className={classes.content}>
          <Typography component="h5" variant="h5">
            Unlucky! no character found, try again.
          </Typography>
        </CardContent>
        <div className={classes.controls}>
          <IconButton aria-label="retry" onClick={props.retry}>
            <CasinoIcon className={classes.retryIcon} />
          </IconButton>
        </div>
      </div>
      <CardMedia className={classes.cover} image={sadMorty} title="Sad Morty" />
    </Card>
  );
};

NotFoundCard.propTypes = { retry: PropTypes.func };

export default NotFoundCard;
