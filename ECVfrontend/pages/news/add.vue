<template>
  <div>
    <v-form>
      <v-container>
        <v-row>
          <v-col>
            <v-text-field
              v-model="newsForm.form.title"
              :counter="50"
              :rules="[v=>v.length > 0 || 'Title is required']"
              label="Title"
              required
            ></v-text-field>
          </v-col>
          <v-col cols="2">
            <v-select
              v-model="newsForm.form.category"
              :items="newsForm.categories"
              :rules="[v => v.length > 0 || 'Category is required']"
              label="Category"
              required
            ></v-select>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <client-only>
              <mavon-editor
                v-model="newsForm.form.text"
                :disabled="newsForm.loading||newsForm.ok"
                style="z-index:0"
              />
            </client-only>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-timeline>
              <v-timeline-item
                v-for="(event, i) in newsForm.form.events"
                :key="i"
                :color="newsForm.colors[i%newsForm.colors.length]"
                small
              >
                <template v-slot:opposite>
                  <span
                    :class="`headline font-weight-bold ${newsForm.colors[i%newsForm.colors.length]}--text`"
                    v-text="event.date"
                  ></span>
                </template>
                <div class="py-2">
                  <h2 :class="`headline font-weight-light mb-2 ${newsForm.colors[i%newsForm.colors.length]}--text`">{{event.addr}}</h2>
                  <div>
                    {{event.description}}
                  </div>
                </div>
              </v-timeline-item>
            </v-timeline>
          </v-col>
        </v-row>

      </v-container>
    </v-form>
    <v-dialog
      max-width="60%"
      scrollable
    >
      <template v-slot:activator="{ on }">
        <v-btn
          dark
          fab
          bottom
          right
          fixed
          x-large
          color="primary"
          v-on="on"
        >
          <v-icon>mdi-plus</v-icon>
        </v-btn>

      </template>
      <v-card>
        <v-card-title>Add Event</v-card-title>
        <v-divider></v-divider>
        <v-card-text>
          <v-form>
            <v-container>
              <v-row>
                <v-col>
                  <v-text-field
                    v-model="eventForm.addr"
                    :counter="50"
                    :rules="[v=>v.length > 0 || 'Address is required']"
                    label="Address"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col>
                  <v-menu
                    :close-on-content-click="false"
                    :nudge-right="40"
                    transition="scale-transition"
                    offset-y
                  >
                    <template v-slot:activator="{ on }">
                      <v-text-field
                        v-model="eventForm.startDate"
                        class="mt-1"
                        label="Start Date"
                        prepend-icon="mdi-calendar"
                        dense
                        readonly
                        outlined
                        hide-details
                        v-on="on"
                      ></v-text-field>
                    </template>
                    <v-date-picker
                      v-model="eventForm.startDate"
                      scrollable
                      dark
                    >
                      <v-spacer></v-spacer>
                      <v-btn
                        text
                        color="primary"
                      >
                        Cancel
                      </v-btn>
                      <v-btn
                        text
                        color="primary"
                      >
                        OK
                      </v-btn>
                    </v-date-picker>
                  </v-menu>
                </v-col>
                <v-col>
                  <v-menu
                    :close-on-content-click="false"
                    :nudge-right="40"
                    transition="scale-transition"
                    offset-y
                  >
                    <template v-slot:activator="{ on }">
                      <v-text-field
                        v-model="eventForm.startTime"
                        class="mt-1"
                        label="Start Time"
                        dense
                        readonly
                        outlined
                        hide-details
                        v-on="on"
                      ></v-text-field>
                    </template>
                    <v-time-picker
                      v-model="eventForm.startTime"
                      format="24hr"
                      dark
                    >
                      <v-spacer></v-spacer>
                      <v-btn
                        text
                        color="primary"
                      >
                        Cancel
                      </v-btn>
                      <v-btn
                        text
                        color="primary"
                      >
                        OK
                      </v-btn>
                    </v-time-picker>
                  </v-menu>
                </v-col>
              </v-row>
              <v-row>
                <v-col>
                  <v-menu
                    :close-on-content-click="false"
                    :nudge-right="40"
                    transition="scale-transition"
                    offset-y
                  >
                    <template v-slot:activator="{ on }">
                      <v-text-field
                        v-model="eventForm.endDate"
                        class="mt-1"
                        label="End Date"
                        prepend-icon="mdi-calendar"
                        dense
                        readonly
                        outlined
                        hide-details
                        v-on="on"
                      ></v-text-field>
                    </template>
                    <v-date-picker
                      v-model="eventForm.endDate"
                      scrollable
                      dark
                    >
                      <v-spacer></v-spacer>
                      <v-btn
                        text
                        color="primary"
                      >
                        Cancel
                      </v-btn>
                      <v-btn
                        text
                        color="primary"
                      >
                        OK
                      </v-btn>
                    </v-date-picker>
                  </v-menu>
                </v-col>
                <v-col>
                  <v-menu
                    :close-on-content-click="false"
                    :nudge-right="40"
                    transition="scale-transition"
                    offset-y
                  >
                    <template v-slot:activator="{ on }">
                      <v-text-field
                        v-model="eventForm.endTime"
                        class="mt-1"
                        label="End Time"
                        dense
                        readonly
                        outlined
                        hide-details
                        v-on="on"
                      ></v-text-field>
                    </template>
                    <v-time-picker
                      v-model="eventForm.endTime"
                      format="24hr"
                      dark
                    >
                      <v-spacer></v-spacer>
                      <v-btn
                        text
                        color="primary"
                      >
                        Cancel
                      </v-btn>
                      <v-btn
                        text
                        color="primary"
                      >
                        OK
                      </v-btn>
                    </v-time-picker>
                  </v-menu>
                </v-col>
              </v-row>
              <v-row>
                <v-col>
                  <v-textarea
                    outlined
                    name="description"
                    label="Description"
                    :value="eventForm.description"
                  ></v-textarea>
                </v-col>
              </v-row>
            </v-container>
          </v-form>
        </v-card-text>

      </v-card>
    </v-dialog>
  </div>
</template>
<script>
import Vue from 'vue'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
Vue.use(mavonEditor)

export default {
  auth: "guest",
  data: function () {
    return {
      newsForm: {
        form: {
          title: "",
          text: "",
          category: "",
          events: [
            {
              date: "2020-01-22",
              addr: "Tianjin No.1 High School",
              description: "Lorem ipsum dolor sit amet, no nam oblique veritus. Commune scaevola imperdiet nec ut, sed euismod convenire principes at. Est et nobis iisque percipit, an vim zril disputando voluptatibus, vix an salutandi sententiae."
            }
          ],
        },
        loading: false,
        ok: false,
        error: false,
        categories: [
          "Tianjin",
          "Beijing"
        ],
        colors: [
          'cyan', 'green', 'pink', 'amber', 'orange'
        ]
      },
      eventForm: {
        startDate: "",
        startTime: "",
        endDate: "",
        endTime: "",
        addr: "",
        description: "",
      }
    }
  }
}
</script>
