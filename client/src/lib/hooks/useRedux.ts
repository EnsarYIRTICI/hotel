import { ref, onUnmounted } from "vue";
import store from "../store";

export const useRedux = (selector) => {
  const state = ref(selector(store.getState())); // Initialize with the state selector

  const updateState = () => {
    state.value = selector(store.getState()); // Update the state on store change
  };

  // Subscribe to the Redux store and update state on changes
  store.subscribe(updateState);

  // Clean up the subscription when the component is unmounted
  onUnmounted(() => {
    store.subscribe(updateState);
  });

  return state;
};

export const useDispatch = () => {
  return store.dispatch;
};
