import { Card, CardContent, Grid, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import CustomBreadcrumbs from "@/components/breadcrumbs";
import { fetchTestOne, getCurrentTest, updateTest } from "@/store/admin/test";
import TestForm from "@/components/form/test/form";
import { testValues } from "@/@local/table/form-values/test/defaultValues";

const TestEdit = () => {
  const [values, setValues] = useState(testValues);
  const [loading, setLoading] = useState(true);

  const dispatch = useDispatch();
  const router = useRouter();
  const { test } = router.query;

  const testValue = useSelector(getCurrentTest);

  useEffect(() => {
    if (test) {
      dispatch(fetchTestOne({ id: test }));
    }
  }, [test, dispatch]);

  useEffect(() => {
    if (testValue && testValue.id) {
      setValues({
        input: testValue.input || "",
        output: testValue.output || "",
      });
      setLoading(false);
    }
  }, [testValue]);

  const handleSubmit = (dataNew) => {
    const data = {
      inputValue: dataNew.input,
      outputValue: dataNew.output,
    };
    dispatch(
      updateTest({
        id: testValue.id,
        ...data,
        callback: () => router.replace(`/admin/chapter/${chapterId}/test`),
      })
    );
  };

  if (loading) {
    return <Typography>Loading...</Typography>;
  }

  return (
    <Grid container spacing={2}>
      <Grid item xs={12} md={12}>
        <CustomBreadcrumbs
          titles={[
            { title: "Admin", path: "/admin" },
            { title: "Courses", path: "/admin/courses" },
            {
              title: "Chapters",
              path: `/admin/courses/${testValue?.chapterID}/chapters`,
            },
            { title: "Edit Test" },
          ]}
        />
        <Typography variant="h2" sx={{ mt: 2 }}>
          Edit Test
        </Typography>
      </Grid>

      <Grid item xs={12} md={12}>
        <Card>
          <CardContent>
            <TestForm
              values={values}
              setValues={setValues}
              handleSubmit={handleSubmit}
              isEdit={true}
            />
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
};

export default TestEdit;
