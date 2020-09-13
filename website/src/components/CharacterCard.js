import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import CardContent from '@material-ui/core/CardContent';
import { PropTypes } from 'prop-types';

const useStyles = makeStyles({
  root: {
    width: 200,
    height: 300,
    margin: 5,
  },
});

const CharacterCard = (props) => {
  const classes = useStyles();
  const { image, name } = props;

  return (
    <Card className={classes.root}>
      <CardActionArea>
        <CardMedia
          component="img"
          alt="Ricky and Morty image"
          image={image}
          title={name}
        />
        <CardContent>
          <Typography gutterBottom component="h3">
            {name}
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
};

CharacterCard.propTypes = { image: PropTypes.string, name: PropTypes.string };

export default CharacterCard;
