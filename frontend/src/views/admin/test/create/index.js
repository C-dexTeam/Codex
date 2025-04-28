import { useRouter } from "next/router";
import React from "react";
import { useDispatch } from "react-redux";
import AttributeForm from "@/components/form/attribute/form";
import { Card, CardContent, Grid, Typography } from "@mui/material";
import { useState } from "react";
import CustomBreadcrumbs from "@/components/breadcrumbs";
import { testValues } from "@/@local/table/form-values/test/defaultValues";
import { createTest } from "@/store/admin/test";
import TestForm from "@/components/form/test/form";

const TetsCreate = () => {
  const dispatch = useDispatch();
  const router = useRouter();
  const { chapterId } = router.query;

  // ** State
  const [values, setValues] = useState(testValues);

  // ** Handlers
  const handleSubmit = (dataNew) => {
    const data = {
      inputValue: dataNew.input,
      outputValue: dataNew.output,
      chapterID: chapterId,
    };
    dispatch(
      createTest({
        data,
          callback: () => router.replace(`/admin/chapter/${chapterId}/test`),
      })
    );
  };

  return (
    <Grid container spacing={2}>
      <Grid item xs={12} md={12}>
        <CustomBreadcrumbs
          titles={[
            { title: "Admin", path: "/admin" },
            { title: "Create Test" },
          ]}
        />
        <Typography variant="h2" sx={{ mt: 2 }}>
          Create Test
        </Typography>
      </Grid>

      <Grid item xs={12} md={12}>
        <Card>
          <CardContent>
            <TestForm
              values={values}
              setValues={setValues}
              handleSubmit={(data) => handleSubmit(data)}
              isEdit={false}
            />
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
};

export default TetsCreate;
