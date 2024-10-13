import { createApi } from "@reduxjs/toolkit/query/react";
import { axiosBaseQuery } from "@app/api/baseQuery";
import { identityAPIHandler } from "@app/api/handlers";

interface Response<T> {
  data: T,
  message: string,
}

interface Project {
  id: string;
  name: string;
  owner_id: number;
  created_at: string;
  owner: {
    name: string;
    email: string;
  }
}

const projectAPI = createApi({
  reducerPath: "projectAPI",
  baseQuery: axiosBaseQuery(identityAPIHandler),
  tagTypes: ["projects"],
  endpoints: (builder) => ({
    getAllProjects: builder.query<Response<Project[]>, void>({
      query: () => ({
        url: "projects/all",
        method: "get",
      }),
      providesTags: ["projects"],
    }),
    createProject: builder.mutation<
      Response<Project>,
      { name: string; }
    >({
      query: ({ name }) => ({
        url: "projects",
        method: "post",
        data: { name },
      }),
      invalidatesTags: ["projects"],
    }),
  }),
});

export const { useGetAllProjectsQuery, useCreateProjectMutation } = projectAPI;

export default projectAPI;
