<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <el-row>
          <el-col :span="12" :offset="6">
            <h1>gin-gorm-vuejs</h1>
          </el-col>
        </el-row>
      </el-header>
      <el-main>
        <el-row>
          <el-col :span="12" :offset="6">
            <el-form
              ref="formRef"
              :model="myForm"
              label-width="100px"
              class="demo-ruleForm"
            >
              <el-form-item label="用户名" prop="username">
                <el-input
                  v-model.string="myForm.username"
                  type="text"
                  autocomplete="off"
                />
              </el-form-item>

              <el-form-item label="密码" prop="password">
                <el-input
                  v-model.string="myForm.password"
                  type="password"
                  autocomplete="off"
                />
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="submitForm(formRef)"
                  >添加</el-button
                >
                <el-button @click="resetForm(formRef)">重置</el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="12" :offset="6">
            <el-table :data="tableData" style="width: 100%">
              <el-table-column prop="username" label="用户名" width="180" />
              <el-table-column prop="password" label="密码" />
            </el-table>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>



<script setup lang="ts">
import { reactive, ref, onMounted } from "vue";
import type { FormInstance } from "element-plus";

import instance from "./axios/index";

let tableData = ref([]);

const addUsers = async () => {
  instance
    .post("http://localhost:8080/add", {
      username: myForm.username,
      password: myForm.password,
    })
    .then((res) => {
      console.log(res.data);
      getUsers();
    });
};

const getUsers = async () => {
  instance.get("http://localhost:8080/users").then((res) => {
    tableData.value = res.data.data;
    console.log(tableData);
  });
};

onMounted(() => {
  getUsers();
});

const formRef = ref<FormInstance>();

const myForm = reactive({
  username: "",
  password: "",
});

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  addUsers();
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};
</script>



<style>
h1 {
  text-align: center;
}
</style>
