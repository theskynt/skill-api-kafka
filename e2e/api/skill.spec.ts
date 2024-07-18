import { test, expect } from "@playwright/test";

test("should response one skill when request /api/v1/skills/:key", async ({
  request,
}) => {
  const reps = await request.get("/api/v1/skills/go");

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "go",
        name: "Go",
        description: expect.any(String),
        logo: expect.any(String),
        tags: expect.arrayContaining(["go", "golang"]),
      },
    })
  );
});

test("should response error when request /api/v1/skills/:key with not exits key", async ({
  request,
}) => {
  const reps = await request.get("/api/v1/skills/go2");

  expect(reps.status() == 500).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: "error",
      message: "skill not found",
    })
  );
});

test("should create a new skill when request POST /api/v1/skills", async ({
  request,
}) => {
  const skillData = {
    key: "js",
    name: "JavaScript",
    description: "A versatile programming language",
    logo: "js-logo.png",
    tags: ["js", "javascript"],
  };

  const reps = await request.post("/api/v1/skills", { data: skillData });

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: skillData,
    })
  );
});

test("should update a skill when request PUT /api/v1/skills/:key", async ({
  request,
}) => {
  const skillData = {
    key: "js",
    name: "JavaScript",
    description: "A popular programming language",
    logo: "js-logo-updated.png",
    tags: ["js", "javascript", "web"],
  };

  const reps = await request.put("/api/v1/skills/js", { data: skillData });

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: skillData,
    })
  );
});

test("should update skill name when request PATCH /api/v1/skills/:key/actions/name", async ({
  request,
}) => {
  const skillNameUpdate = { name: "New JavaScript" };

  const reps = await request.patch("/api/v1/skills/js/actions/name", {
    data: skillNameUpdate,
  });

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "js",
        name: "New JavaScript",
        description: expect.any(String),
        logo: expect.any(String),
        tags: expect.arrayContaining(["js", "javascript"]),
      },
    })
  );
});

test("should update skill description when request PATCH /api/v1/skills/:key/actions/description", async ({
  request,
}) => {
  const skillDescriptionUpdate = { description: "Updated description" };

  const reps = await request.patch("/api/v1/skills/js/actions/description", {
    data: skillDescriptionUpdate,
  });

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "js",
        name: expect.any(String),
        description: "Updated description",
        logo: expect.any(String),
        tags: expect.arrayContaining(["js", "javascript"]),
      },
    })
  );
});

test("should update skill logo when request PATCH /api/v1/skills/:key/actions/logo", async ({
  request,
}) => {
  const skillLogoUpdate = { logo: "new-logo.png" };

  const reps = await request.patch("/api/v1/skills/js/actions/logo", {
    data: skillLogoUpdate,
  });

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "js",
        name: expect.any(String),
        description: expect.any(String),
        logo: "new-logo.png",
        tags: expect.arrayContaining(["js", "javascript"]),
      },
    })
  );
});

test("should update skill tags when request PATCH /api/v1/skills/:key/actions/tags", async ({
  request,
}) => {
  const skillTagsUpdate = { tags: ["updated", "tags"] };

  const reps = await request.patch("/api/v1/skills/js/actions/tags", {
    data: skillTagsUpdate,
  });

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "js",
        name: expect.any(String),
        description: expect.any(String),
        logo: expect.any(String),
        tags: ["updated", "tags"],
      },
    })
  );
});

test("should delete a skill when request DELETE /api/v1/skills/:key", async ({
  request,
}) => {
  const reps = await request.delete("/api/v1/skills/js");

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: "Skill deleted",
    })
  );
});
