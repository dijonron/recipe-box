import React, { useState } from "react";
import AddIcon from "@mui/icons-material/Add";
import {
  Button,
  CircularProgress,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TablePagination,
  TableRow,
  Typography,
} from "@mui/material";
import useFetch from "../../hooks/useFetch";
import { RecipeModal } from "../RecipeModal";
import DeleteModal from "../DeleteModal/DeleteModal.component";
import { useSnackbar } from "../../contexts/SnackbarContext";

function Recipies() {
  const [recipeModal, setRecipeModal] = useState(false);
  const [deleteModal, setDeleteModal] = useState(false);

  const { showSnackbar } = useSnackbar();

  const count = 15;
  const rowsPerPage = 5;
  const page = 0;

  const handleChangePage = () => {};
  const handleChangeRowsPerPage = () => {};

  const { data, error, loading, refetch } = useFetch(
    `${process.env.REACT_APP_GATEWAY_URL}/v1/recipes`,
    {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    }
  );

  if (error) {
    return <>Error!</>;
  }

  const { recipes = [] } = data ?? {};
  const renderLoading = loading;
  const renderEmpty = !recipes.length;

  const deleteRecipe = async ({ id }) => {
    try {
      const response = await fetch(
        `${process.env.REACT_APP_GATEWAY_URL}/v1/recipes/${id}`,
        {
          method: "DELETE",
        }
      );

      if (!response.ok) {
        throw new Error(`Error: ${response.statusText}`);
      }
      await refetch();
      showSnackbar("Recipe deleted!", "success");
    } catch (error) {
      showSnackbar(`Unable to delete the recipe. ${error.message}`, "error");
    } finally {
      setDeleteModal(false);
    }
  };

  return (
    <>
      <Button
        size="medium"
        color="primary"
        variant="contained"
        startIcon={<AddIcon />}
        onClick={() => {
          setRecipeModal(true);
        }}
      >
        Add Recipie
      </Button>
      {recipeModal && (
        <RecipeModal
          open={recipeModal}
          handleClose={async () => {
            await refetch();
            setRecipeModal(false);
          }}
        />
      )}
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Recipie</TableCell>
              <TableCell>Chef</TableCell>
              <TableCell></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {renderLoading && (
              <TableRow>
                <TableCell colSpan={3} align="center">
                  <CircularProgress size="3rem" />
                </TableCell>
              </TableRow>
            )}
            {!renderLoading && renderEmpty ? (
              <TableRow>
                <TableCell colSpan={3}>
                  <Typography align="center">
                    Please add some recipes!
                  </Typography>
                </TableCell>
              </TableRow>
            ) : (
              recipes?.map((recipe) => (
                <TableRow key={`recipe-row--${recipe.id}`}>
                  <TableCell>{recipe.name}</TableCell>
                  <TableCell>{recipe.chef}</TableCell>
                  <TableCell align="right">
                    <Button size="small" color="secondary" variant="outlined">
                      Edit
                    </Button>
                    <span style={{ marginLeft: "16px" }}> </span>
                    <Button
                      size="small"
                      color="error"
                      variant="text"
                      onClick={() => {
                        setDeleteModal(recipe);
                      }}
                    >
                      Delete
                    </Button>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[5, 10, 25]}
        component="div"
        count={count}
        rowsPerPage={rowsPerPage}
        page={page}
        onPageChange={handleChangePage}
        onRowsPerPageChange={handleChangeRowsPerPage}
      />
      {deleteModal && (
        <DeleteModal
          recipe={deleteModal}
          open={!!deleteModal}
          handleClose={() => {
            setDeleteModal(false);
          }}
          handleSubmit={() => deleteRecipe(deleteModal)}
        />
      )}
    </>
  );
}

export default Recipies;
