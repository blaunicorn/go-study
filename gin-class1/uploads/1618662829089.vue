<template>
  <div>
    <el-dialog v-bind="$attrs" v-on="$listeners" @open="onOpen" @close="onClose" title="Dialog Titile">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
        <el-form-item label="手机号" prop="mobile">
          <el-input v-model="formData.mobile" placeholder="请输入手机号" :maxlength="11" show-word-limit clearable
            prefix-icon='el-icon-mobile' :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item label="级联选择" prop="field101">
          <el-cascader v-model="formData.field101" :options="field101Options" :props="field101Props"
            :style="{width: '100%'}" placeholder="请选择级联选择" clearable></el-cascader>
        </el-form-item>
      </el-form>
      <div slot="footer">
        <el-button @click="close">取消</el-button>
        <el-button type="primary" @click="handelConfirm">确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
export default {
  inheritAttrs: false,
  components: {},
  props: [],
  data() {
    return {
      formData: {
        mobile: '',
        field101: [],
      },
      rules: {
        mobile: [{
          required: true,
          message: '请输入手机号',
          trigger: 'blur'
        }, {
          pattern: /^1(3|4|5|7|8|9)\d{9}$/,
          message: '手机号格式错误',
          trigger: 'blur'
        }],
        field101: [{
          required: true,
          type: 'array',
          message: '请至少选择一个field101',
          trigger: 'change'
        }],
      },
      field101Options: [],
      field101Props: {
        "multiple": false
      },
    }
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {},
  methods: {
    onOpen() {},
    onClose() {
      this.$refs['elForm'].resetFields()
    },
    close() {
      this.$emit('update:visible', false)
    },
    handelConfirm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        this.close()
      })
    },
    getField101Options() {
      // TODO 发起请求获取数据
      this.field101Options
    },
  }
}

</script>
<style>
</style>
