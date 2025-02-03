import React, { Fragment, useEffect, useState } from "react";
import { useFormik } from "formik";
import AddIcon from "@mui/icons-material/Add";
import CloseIcon from "@mui/icons-material/Close";
import {
  Autocomplete,
  Button,
  Dialog,
  DialogContent,
  DialogTitle,
  FormControl,
  Grid2 as Grid,
  InputLabel,
  MenuItem,
  Select,
  TextField,
  Typography,
} from "@mui/material";
import { MEASURMENT } from "../../utils/constants";

function RecipeModal(props) {
  const { open = false, handleClose } = props;

  const [ingredients, setIngredients] = useState([]);

  const formik = useFormik({
    initialValues: {
      name: "",
      chef: "",
      cookbook: "",
      ingredients: [
        {
          ingredient_id: "",
          name: "",
          amount: 1,
          measurement: "",
        },
        {
          ingredient_id: "",
          name: "",
          amount: 1,
          measurement: "",
        },
        {
          ingredient_id: "",
          name: "",
          amount: 1,
          measurement: "",
        },
      ],
    },
    onSubmit: async (values) => {
      const res = await fetch(
        `${process.env.REACT_APP_GATEWAY_URL}/v1/recipes`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ recipe: values }),
        }
      );

      const data = await res.json();
      console.log(data);
    },
  });

  useEffect(() => {
    const fetchIngredients = async () => {
      const res = await fetch(
        `${process.env.REACT_APP_GATEWAY_URL}/v1/ingredients`,
        {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        }
      );

      if (!res.ok) {
        throw new Error("Network error");
      }

      const { ingredients } = await res.json();

      setIngredients(ingredients);
    };

    fetchIngredients();
  }, []);

  return (
    <Dialog open={open} onClose={handleClose}>
      <DialogTitle>Fill in the recipe's details</DialogTitle>
      <DialogContent>
        <form onSubmit={formik.handleSubmit}>
          <Grid container spacing={2}>
            <Grid size={12}>
              <TextField
                id="name"
                name="name"
                label="Recipe"
                variant="standard"
                fullWidth
                value={formik.values.name}
                onChange={formik.handleChange}
                onBlur={formik.handleBlur}
                error={formik.touched.name && Boolean(formik.errors.name)}
                helperText={formik.touched.name && formik.errors.name}
              />
            </Grid>
            <Grid size={6}>
              <TextField
                id="chef"
                name="chef"
                label="Chef"
                variant="standard"
                fullWidth
                value={formik.values.chef}
                onChange={formik.handleChange}
                onBlur={formik.handleBlur}
                error={formik.touched.chef && Boolean(formik.errors.chef)}
                helperText={formik.touched.chef && formik.errors.chef}
              />
            </Grid>
            <Grid size={6}>
              <TextField
                id="cookbook"
                name="cookbook"
                label="Cookbook"
                variant="standard"
                fullWidth
                value={formik.values.cookbook}
                onChange={formik.handleChange}
                onBlur={formik.handleBlur}
                error={
                  formik.touched.cookbook && Boolean(formik.errors.cookbook)
                }
                helperText={formik.touched.cookbook && formik.errors.cookbook}
              />
            </Grid>

            <Grid size={12}>
              <Typography variant="body2">Ingredients</Typography>
            </Grid>
          </Grid>

          <Grid container spacing={2} alignItems={"flex-end"}>
            {formik.values.ingredients.map((ingredient, i) => {
              return (
                <Fragment key={`ingredient-${i}`}>
                  <Grid size={6}>
                    <Autocomplete
                      disablePortal
                      freeSolo
                      filterSelectedOptions
                      options={ingredients}
                      id={`ingredients-${i}-name`}
                      name={`ingredients.${i}.name`}
                      value={formik.values.ingredients[i]}
                      onChange={(_, newValue) => {
                        formik.setFieldValue(
                          `ingredients.${i}.name`,
                          newValue?.name ?? ""
                        );
                        formik.setFieldValue(
                          `ingredients.${i}.ingredient_id`,
                          newValue?.id ?? ""
                        );
                      }}
                      onInputChange={(_, newValue) => {
                        formik.setFieldValue(
                          `ingredients.${i}.name`,
                          newValue ?? ""
                        );
                      }}
                      onBlur={formik.handleBlur}
                      getOptionLabel={({ name }) => name}
                      getOptionKey={({ ingredient_id }) => ingredient_id}
                      renderInput={(params) => {
                        return (
                          <TextField
                            {...params}
                            label="Ingredient"
                            variant="standard"
                            fullWidth
                          />
                        );
                      }}
                    />
                  </Grid>
                  <Grid size={2}>
                    <TextField
                      id={`ingredient-${i}-amount`}
                      name={`ingredients.${i}.amount`}
                      label="Amount"
                      variant="standard"
                      type="number"
                      fullWidth
                      value={formik.values.ingredients[i].amount}
                      onChange={formik.handleChange}
                      onBlur={formik.handleBlur}
                    />
                  </Grid>
                  <Grid size={3}>
                    <FormControl variant="standard" fullWidth>
                      <InputLabel id="measurement-select-label">
                        Unit
                      </InputLabel>
                      <Select
                        id={`ingredient-${i}-measurement`}
                        name={`ingredients.${i}.measurement`}
                        label="Unit"
                        variant="standard"
                        fullWidth
                        value={formik.values.ingredients[i].measurement}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                      >
                        {Object.entries(MEASURMENT).map(([val, key]) => (
                          <MenuItem value={val}>{key}</MenuItem>
                        ))}
                      </Select>
                    </FormControl>
                  </Grid>
                  <Grid size={1}>
                    <Button
                      size="medium"
                      color="secondary"
                      variant="text"
                      onClick={() => {
                        formik.setFieldValue("ingredients", [
                          ...formik.values.ingredients.slice(0, i),
                          ...formik.values.ingredients.slice(i + 1),
                        ]);
                      }}
                      startIcon={<CloseIcon />}
                    />
                  </Grid>
                </Fragment>
              );
            })}
            <Grid size={12}>
              <Button
                size="medium"
                color="secondary"
                variant="text"
                onClick={() => {
                  formik.setFieldValue("ingredients", [
                    ...formik.values.ingredients,
                    { id: 0, name: "", amount: 1, measurement: "" },
                  ]);
                }}
                startIcon={<AddIcon />}
              >
                Add Ingredient
              </Button>
            </Grid>
          </Grid>

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
              type="submit"
              // onClick={() => console.log("add recipe")}
            >
              Add Recipie
            </Button>
          </Grid>
        </form>
      </DialogContent>
    </Dialog>
  );
}

export default RecipeModal;
