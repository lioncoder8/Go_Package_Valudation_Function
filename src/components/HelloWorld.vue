<template>
  <v-row justify="center">
    <v-dialog
      v-model="dialogEdit"
      max-width="900"
      :content-class="'dialogEdit'"
      persistent
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          color="primary"
          dark
          v-bind="attrs"
          v-on="on"
        >
          Open Dialog
        </v-btn>
      </template>
      <v-card>
        <v-form ref="form" v-model="createDisplay">
          <v-card-title class="edit-close">
            <span @click="closeItemEditDialog()">
              <icon-delete/>
            </span>
          </v-card-title>
          <v-card-title class="title-add">
            項目設定
            <v-select
              v-model="selectedGroup"
              :items="groupInfo"
              label="事業部選択"
              item-text="name"
              item-value="id"
              class="group-select"
              @change="selectGroup($event)"
              solo
            >
            </v-select>
            <v-btn class="btn-submit">保存</v-btn>
          </v-card-title>
          <div>
            <div class="sidenav">
              <div class="sidenav-item">項目名</div>
              <div class="sidenav-item">入力方法</div>
              <div class="sidenav-item">選択項目</div>
            </div>
            <div class="sidebar">
              <div class="item">
                <div class="item-0">
                  <span class="icon-order">
                    <span class="icon-left">
                      <icon-left/>
                    </span>
                    <span class="icon-right">
                      <icon-right/>
                    </span>
                  </span>
                  <span class="delete-item">
                      <icon-delete/><br>
                  </span>
                </div>
                <div class="item-1">
                  <input
                    type="text"
                    placeholder="項目名"
                    class="form-control"
                  />
                  <span>
                    <icon-base icon-color="#14659A" height="15" width="15">
                      <icon-plus/>
                    </icon-base>
                  </span>
                </div>
                <div class="item-2">
                  <v-select
                    :items="inputMethods"
                    label="入力方法"
                    item-text="text"
                    item-value="value"
                    solo
                  >
                  </v-select>
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    placeholder="選択項目名"
                    >
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    placeholder="選択項目名"
                    >
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    placeholder="選択項目名"
                    >
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    placeholder="選択項目名"
                    >
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    style="margin-bottom: 10px"
                    placeholder="選択項目名1"
                    >
                </div>
              </div>
              <div class="item">
                <div class="item-0">
                  <span class="icon-order">
                    <span class="icon-left">
                      <icon-left/>
                    </span>
                    <span class="icon-right">
                      <icon-right/>
                    </span>
                  </span>
                  <span class="delete-item">
                      <icon-delete/><br>
                  </span>
                </div>
                <div class="item-1">
                  <input
                    type="text"
                    placeholder="項目名"
                    class="form-control"
                  />
                  <span>
                    <icon-base icon-color="#14659A" height="15" width="15">
                      <icon-plus/>
                    </icon-base>
                  </span>
                </div>
                <div class="item-2">
                  <v-select
                    :items="inputMethods"
                    label="入力方法"
                    item-text="text"
                    item-value="value"
                    solo
                  >
                  </v-select>
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    placeholder="選択項目名"
                    >
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    placeholder="選択項目名"
                    >
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    placeholder="選択項目名"
                    >
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    placeholder="選択項目名"
                    >
                </div>
                <div class="item-3">
                  <input
                    class="form-control select-item"
                    style="margin-bottom: 10px"
                    placeholder="選択項目名1"
                    >
                </div>
              </div>
            </div>
          </div>
        </v-form>
      </v-card>
    </v-dialog>
  </v-row>
</template>
<script>
import IconDelete from './IconDelete.vue'
import IconLeft from './iconLeft.vue'
import IconRight from '../components/iconRight.vue'
import IconPlus from '../components/IconPlus.vue'
export default {
  components: {
    IconDelete,
    IconLeft,
    IconRight,
    IconPlus,
  },
  data(){
    return {
      dialogEdit: false,
      groupInfo: ["保存保存1", "保存保存", "保存"],
      inputMethods: ["保存保存1", "保存保存", "保存"]
    }
  },
  methods: {
    closeItemEditDialog() {
      this.dialogEdit = false
    }
  }
}
</script>
<style>
.sidenav {
  width: 140px;
  position: fixed;
  z-index: 1;
  overflow-x: hidden;
  padding-top: 50px;
  padding-left: 45px;
}
.sidenav-item {
  margin-bottom: 40px;
}
.sidebar {
  margin-left: 122px;
  margin-right: 45px;
  display: flex;
  overflow-x: scroll;
  padding-top: 16px;
  height: 405px;
}
.dialogEdit .sidebar .icon-order {
  vertical-align: text-bottom;
}
.dialogEdit .sidebar .delete-item {
  margin-left: 110px;
}
.dialogEdit .sidebar .delete-item svg {
  width: 15px;
}
.dialogEdit .item .icon-order {
  padding-left: 20px;
}
.dialogEdit .item svg {
  cursor: pointer;
  pointer-events: auto;
}
.dialogEdit .item .item-1 {
  padding-bottom: 15px;
  border-bottom: 2px dotted #E6E6E6;
}
.dialogEdit .item .item-2 {
  padding: 15px 0;
  border-bottom: 2px dotted #E6E6E6;
}
.dialogEdit .item .item-3 {
  padding: 15px 0 0 0;
}
.dialogEdit .item .form-control {
  border: 1px solid #E6E6E6;
  border-radius: 5px;
  padding-left: 10px;
  margin-left: 0px;
  width: 200px;
  height: 35px;
}
.dialogEdit .item .v-select .v-input__slot {
  margin-left: 0;
  border: 1px solid white;
}


.dialogEdit {
  border-radius: 20px;
}
.dialogEdit .v-sheet.v-card {
  min-height: 500px !important;
}
.dialogEdit .edit-close {
  flex-direction: row-reverse;
  padding: 8px 16px 0px;
}
.dialogEdit .edit-close span {
  max-height: 25px;
}
.dialogEdit svg {
  cursor: pointer;
}
.dialogEdit .row {
  margin: 0;
}
.dialogEdit .row svg {
  cursor: pointer;
  pointer-events: auto;
}
.dialogEdit .row .form-control {
  border: 1px solid #E6E6E6;
  border-radius: 5px;
  padding-left: 10px;
  margin-left: 20px;
  width: 210px;
}
.dialogEdit .row .select-item {
  margin-left: 5px;
}
.dialogEdit .v-card form .v-card__title.close {
  font-size: 24px !important;
  text-align: center !important;
  max-width: 30px !important;
  max-height: 30px !important;
  position: absolute !important;
  right: 0 !important;
  background: #70B6D1 !important;
  border-radius: unset !important;
  padding: 2px 0px 5px 9px !important;
  margin: 0px !important;
  color: #FFFFFF !important;
  position: relative !important;
  margin-left: auto !important;
  cursor: pointer;
}
.dialogEdit.v-dialog .v-card form .v-card__title.title-add
{
  padding: 0px 20px 15px 0px !important;
  color: #14659A !important;
  font: normal normal normal 28px/38px NotoSansJP !important;
  margin: 0 auto !important;
  border-bottom: 1px solid #9F9F9F;
  width: 90% !important;
}
.dialogEdit .v-text-field.v-text-field--solo .v-input__control {
  max-height: 32px;
  min-height: 30px;
}
.dialogEdit .v-card form .v-input__slot {
  width: 210px;
  margin-left: 10px;
  margin-bottom: 0;
  padding: 0 0 0 12px;
}
.dialogEdit .v-card form .group-select {
  width: 250px;
}
.dialogEdit form .btn-submit
{
  width: 150px !important;
  height: 30px !important;
  background-color: #14659A !important;
  font: normal normal normal 12px/20px NotoSansJP !important;
  color: #FFFFFF !important;
  letter-spacing: 3px;
  margin: 0;
}
.dialogEdit .edit-line .icon-order {
  margin-left: 95px;
  vertical-align: text-bottom;
}
.dialogEdit .edit-line .delete-item {
  margin-left: 110px;
}
.dialogEdit .edit-line .delete-item svg {
  width: 15px;
}
.dialogEdit .v-input__append-inner {
  padding-left: 10px !important;
  padding-bottom: 4px;
}
.dialogEdit.v-dialog .v-card form .v-select__slot .v-select__selections {
  border-right: 1px solid #E6E6E6;
}
.dialogEdit .v-input__slot:before {
  border: unset !important;
}
.dialogEdit .edit-line {
  border-bottom: 2px dotted #E6E6E6;
  width: 330px;
}
.dialogEdit .edit-content {
  display: flex;
  padding-top: 5px;
  padding-bottom: 15px;
  margin-top: 10px;
  margin-bottom: 10px;
}
.dialogEdit .edit-list {
  display: flex;
  overflow-x: scroll;
  background-color: #fff;
  width: 90%;
  margin: 20px auto 0;
  max-height: 360px;
}
.dialogEdit .edit-content .v-select .v-input__slot {
  width: 210px;
}
.dialogEdit .edit-content span {
  padding-top: 5px;
  width: 65px;
}
.dialogEdit .input-add {
  margin-left: 185px;
}
.dialogEdit form .btn-submit
{
  width: 150px !important;
  height: 30px !important;
  background-color: #14659A !important;
  font: normal normal normal 12px/20px NotoSansJP !important;
  color: #FFFFFF !important;
  letter-spacing: 3px;
  margin: 0;
}

</style>