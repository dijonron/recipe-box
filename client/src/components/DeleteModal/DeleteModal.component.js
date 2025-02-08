import {
  Button,
  Dialog,
  DialogContent,
  DialogTitle,
  Grid2 as Grid,
  Typography,
} from "@mui/material";

function DeleteModal(props) {
  const { open, handleClose, handleSubmit, recipe } = props;
  const { name } = recipe;

  return (
    <Dialog open={open} onClose={handleClose} scroll="body">
      <DialogTitle>Delete recipe</DialogTitle>
      <DialogContent sx={{ overflowX: "hidden" }}>
        <Typography sx={{ paddingBottom: "32px" }}>
          Are you sure you would like to delete the reipce for {name}?
        </Typography>
        <Grid container justifyContent={"flex-end"}>
          <Button
            size="medium"
            color="error"
            variant="text"
            onClick={handleClose}
          >
            Cancel
          </Button>
          <Button
            size="medium"
            color="primary"
            variant="contained"
            onClick={handleSubmit}
          >
            Delete
          </Button>
        </Grid>
      </DialogContent>
    </Dialog>
  );
}

export default DeleteModal;
